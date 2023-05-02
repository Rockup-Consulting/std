package mid

import (
	"context"
	"log"
	"net/http"

	"github.com/Rockup-Consulting/go_std/x/httpx"
)

func CatchErrAndRespond(l *log.Logger) httpx.Middleware {
	mid := func(h httpx.Handler) httpx.Handler {
		hh := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			if err := h(ctx, w, r); err != nil {
				// if we receive a httpx.Error, simply propogate nil
				if httpx.IsError(err) {
					return nil
				}

				l.Printf("ERROR: %s", err.Error())

				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Error"))
			}

			return nil
		}
		return hh
	}

	return mid
}

func CatchErr(l *log.Logger) httpx.Middleware {
	mid := func(h httpx.Handler) httpx.Handler {
		hh := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			if err := h(ctx, w, r); err != nil {
				// if we receive a httpx.Error, simply propogate nil
				if httpx.IsError(err) {
					return nil
				}

				l.Printf("ERROR: %s", err.Error())
			}

			return nil
		}

		return hh
	}

	return mid
}
