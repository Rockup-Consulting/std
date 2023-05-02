package memcache

import (
	"errors"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/Rockup-Consulting/go_std/x/hashx"
)

type memCache struct {
	mu sync.Mutex
	c  map[string]string
}

func (c *memCache) get(route string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	fileHash, ok := c.c[route]
	if ok {
		return fileHash, true
	}

	return "", false
}

func (c *memCache) set(route string, r io.Reader) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	fileHash, err := hashx.MD5(r)
	if err != nil {
		return "", err
	}

	c.c[route] = fileHash
	return fileHash, nil
}

func newMemCache() memCache {
	return memCache{
		c: map[string]string{},
	}
}

func StaticHandler(fs fs.FS, l *log.Logger) http.HandlerFunc {
	c := newMemCache()
	fileServer := http.FileServer(http.FS(fs))

	// create an http handler
	h := func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/")

		// Get asset etag. If the etag has not yet been created, we first create the etag
		etag, ok := c.get(path)

		// if we haven't cached the asset yet, we should do that first
		if !ok {
			l.Printf("etag not set, hashing file: %s", path)
			file, err := fs.Open(path)

			if err != nil {
				// if the file does not exist, we respond with 404
				if errors.Is(err, os.ErrNotExist) {
					l.Print("404 -> file not found")
					w.Header().Set("Content-Type", "text/plain")
					w.WriteHeader(http.StatusNotFound)
					w.Write([]byte("asset not found"))
					return

					// otherwise something went wrong
				} else {
					l.Printf("internal error: %s", err)
					w.Header().Set("Content-Type", "text/plain")
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("something went wrong!"))
					return
				}
			}

			etag, err = c.set(path, file)
			if err != nil {
				l.Printf("internal error: %s", err)
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("something went wrong!"))
				return
			}
		}

		if etag == r.Header.Get("if-none-match") {
			l.Printf("browser has cached file: %s", path)
			w.WriteHeader(http.StatusNotModified)
			return
		} else {
			w.Header().Set("cache-control", "max-age=60")
			w.Header().Set("etag", etag)

			fileServer.ServeHTTP(w, r)
		}
	}

	return h
}
