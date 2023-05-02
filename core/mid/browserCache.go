package mid

import (
	"context"
	"io"
	"io/fs"
	"log"
	"net/http"

	"strings"
	"sync"

	"github.com/Rockup-Consulting/go_std/x/hashx"
	"github.com/Rockup-Consulting/go_std/x/httpx"
)

// Implementation based on: https://www.youtube.com/watch?v=3XkU_DXcgl0
// Learn about cache control headers: https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control

type browserCache struct {
	mu sync.Mutex
	c  map[string]string
}

func (c *browserCache) get(route string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	etag, ok := c.c[route]
	if ok {
		return etag, true
	}

	return "", false
}

func (c *browserCache) set(route string, r io.Reader) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	hashedVal, err := hashx.MD5(r)
	if err != nil {
		return "", err
	}

	c.c[route] = hashedVal
	return hashedVal, nil
}

func newBrowserCache() browserCache {
	return browserCache{
		c: map[string]string{},
	}
}

// BrowserCache lazily caches static assets in memory
func BrowserCache(l *log.Logger, fs fs.FS) httpx.Middleware {
	c := newBrowserCache()

	mid := func(h httpx.Handler) httpx.Handler {
		hh := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			// Get asset etag. If the etag has not yet been created, we first create the etag
			path := strings.TrimPrefix(r.URL.Path, "/")
			etag, ok := c.get(path)
			if !ok {
				l.Printf("etag not set, hashing file: %s", path)
				file, err := fs.Open(path)
				if err != nil {
					l.Print(err)
					return httpx.Text(w, http.StatusNotFound, "asset not found")
				}

				etag, err = c.set(path, file)
				if err != nil {
					return err
				}
			}

			if etag == r.Header.Get("if-none-match") {
				l.Printf("browser has cached file: %s", path)
				w.WriteHeader(http.StatusNotModified)
				return nil
			} else {
				w.Header().Set("cache-control", "max-age=60")
				w.Header().Set("etag", etag)

				if err := h(ctx, w, r); err != nil {
					return err
				}
				return nil
			}

		}
		return hh
	}

	return mid
}
