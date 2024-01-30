// package mid contains default http handler middlewares
package mid

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/Rockup-Consulting/std/core/web"
)

type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (c *logResponseWriter) WriteHeader(statusCode int) {
	c.statusCode = statusCode
	c.ResponseWriter.WriteHeader(statusCode)
}

func Log(l *log.Logger) web.Middleware {
	return func(h web.Handler) web.Handler {
		return func(w http.ResponseWriter, r *http.Request) error {
			l.Printf("started request %s %s", r.Method, r.URL.Path)

			rw := logResponseWriter{w, http.StatusOK}

			err := h(&rw, r)

			// if an error hasn't been handled by the Error Middleware, we have a problem.
			if err != nil {
				return err
			}

			l.Printf("end request %s status -> %d", r.URL.Path, rw.statusCode)

			return nil
		}
	}
}

func Redirect(l *log.Logger, from string, to string) web.Middleware {
	return func(h web.Handler) web.Handler {
		return func(w http.ResponseWriter, r *http.Request) error {
			if r.Host == from {
				l.Printf("redirecting from %q to %q", r.URL.String(), to+r.URL.Path)
				http.Redirect(w, r, to+r.URL.Path, http.StatusTemporaryRedirect)
				return nil
			}

			return h(w, r)
		}
	}
}

func CatchPanic() web.Middleware {
	return func(h web.Handler) web.Handler {
		return func(w http.ResponseWriter, r *http.Request) (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("PANIC: %s", r)
				}
			}()

			return h(w, r)
		}
	}
}

func CatchErr(l *log.Logger) web.Middleware {
	return func(h web.Handler) web.Handler {
		return func(w http.ResponseWriter, r *http.Request) error {
			if err := h(w, r); err != nil {
				// if we receive a web.Err, simply propogate nil
				if web.IsError(err) {
					return nil
				}

				l.Printf("ERROR: %s", err.Error())

				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Error"))
			}

			return nil
		}
	}
}

// TryGzip checks whether the calling client accepts gzip encoding.
// If it does, the response is passed to a gzip writer, otherwise
// the response just goes to the http.ResponseWriter
func TryGzip(h web.Handler) web.Handler {
	return func(w http.ResponseWriter, r *http.Request) error {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			w.Header().Add("Content-Encoding", "gzip")
			gzw := gzip.NewWriter(w)
			defer gzw.Close()

			gzipRW := gzipResponseWriter{
				ResponseWriter: w,
				GzipWriter:     gzw,
			}

			return h(gzipRW, r)
		} else {
			return h(w, r)
		}
	}
}

// gzipResponseWriter implements io.Writer
type gzipResponseWriter struct {
	http.ResponseWriter
	GzipWriter io.Writer
}

func (g gzipResponseWriter) Write(data []byte) (int, error) {
	return g.GzipWriter.Write(data)
}
