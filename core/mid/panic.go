package mid

import (
	"context"
	"fmt"
	"net/http"

	"runtime/debug"

	"github.com/Rockup-Consulting/go_std/x/httpx"
)

func CatchPanic() httpx.Middleware {
	m := func(h httpx.Handler) httpx.Handler {
		hh := func(ctx context.Context, w http.ResponseWriter, r *http.Request) (err error) {
			defer func() {
				if rec := recover(); rec != nil {
					trace := debug.Stack()
					err = fmt.Errorf("recovered panic [%v] TRACE[%s]", rec, string(trace))
				}
			}()

			return h(ctx, w, r)
		}

		return hh
	}

	return m
}
