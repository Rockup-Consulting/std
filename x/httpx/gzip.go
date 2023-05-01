package httpx

import (
	"compress/gzip"
	"context"
	"io"
	"net/http"
	"strings"
)

// Gzip golang explanation - https://www.youtube.com/watch?v=oSDtwcvU2-o

type gzipResponseWriter struct {
	http.ResponseWriter
	GzipWriter io.Writer
}

func (g gzipResponseWriter) Write(data []byte) (int, error) {
	return g.GzipWriter.Write(data)
}

func acceptsGzip(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Accept-Encoding"), "gzip")
}

func serveGzip(ctx context.Context, w http.ResponseWriter, r *http.Request, f Handler) error {
	w.Header().Add("Content-Encoding", "gzip")
	gzw := gzip.NewWriter(w)
	defer gzw.Close()

	gzipRW := gzipResponseWriter{
		ResponseWriter: w,
		GzipWriter:     gzw,
	}

	return f(ctx, gzipRW, r)
}
