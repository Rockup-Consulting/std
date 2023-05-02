package mid

import (
	"context"
	"log"
	"net/http"

	"github.com/Rockup-Consulting/go_std/x/httpx"
)

type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (c *logResponseWriter) WriteHeader(statusCode int) {
	c.statusCode = statusCode
	c.ResponseWriter.WriteHeader(statusCode)
}

func Log(l *log.Logger) httpx.Middleware {
	return func(h httpx.Handler) httpx.Handler {
		hh := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			l.Printf("started request %s %s", r.Method, r.URL.Path)

			rw := logResponseWriter{w, http.StatusOK}

			err := h(ctx, &rw, r)

			// if an error hasn't been handled by the Error Middleware, we have a problem.
			if err != nil {
				return err
			}

			l.Printf("end request %s status -> %d", r.URL.Path, rw.statusCode)

			return nil
		}

		return hh
	}
}
