package web

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/Rockup-Consulting/std/x/logx"
)

// ====================================================================
// ETAG CACHING

type memCache struct {
	mu sync.Mutex
	c  map[string]string
}

func newMemCache() memCache {
	return memCache{
		c: map[string]string{},
	}
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

	fileHash, err := md5Util(r)
	if err != nil {
		return "", err
	}

	c.c[route] = fileHash
	return fileHash, nil
}

func md5Util(r io.Reader) (string, error) {
	h := md5.New()

	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func StaticHandler(fs fs.FS, l *log.Logger, cacheSeconds int, discardLogs bool) http.HandlerFunc {
	c := newMemCache()
	fileServer := http.FileServer(http.FS(fs))

	// overwrite the logger with one that discards logs
	if discardLogs {
		l = logx.NewDiscard()
	}

	// create an http handler
	h := func(w http.ResponseWriter, r *http.Request) {
		// strip the leading / from the URL path to turn it into
		// a valid file path
		path := strings.TrimPrefix(r.URL.Path, "/")

		// Try to get the asset etag
		etag, ok := c.get(path)

		// if we haven't cached the asset yet, we should do that first
		if !ok {
			l.Printf("etag not set, hashing file: %s", path)

			// try to open the file for the given path
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

			// if the file was opened successfully, we set it in our in-mem cache
			etag, err = c.set(path, file)

			// if that fails, we return an internal server error
			if err != nil {
				l.Printf("internal error: %s", err)
				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("internal server error (500)"))
				return
			}
		}

		// try to get the if-none-match header from the request,
		// if it is available and valid, we respond with a
		// status-304 NotModified. This informs the browser that
		// it may use the asset that is still cached - sending almost
		// zero bytes over the network
		if etag == r.Header.Get("if-none-match") {
			l.Printf("client has cached file: %s", path)
			w.WriteHeader(http.StatusNotModified)
			return

			// if the client’s browser does not have the 'if-none-match'
			// header present in the request, we set the 'cache-control' and
			// 'etag' headers
		} else {
			w.Header().Set("cache-control", fmt.Sprintf("max-age=%d", cacheSeconds))
			w.Header().Set("etag", etag)

			fileServer.ServeHTTP(w, r)
		}
	}

	return h
}
