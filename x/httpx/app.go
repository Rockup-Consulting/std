package httpx

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dimfeld/httptreemux"
)

// A Handler is a type that handles a http request within our own little mini
// framework.
type Handler func(
	ctx context.Context,
	w http.ResponseWriter,
	r *http.Request,
) error

const (
	MegaByte = 1 << 20
)

// App is the entrypoint into our application
type App struct {
	mux *httptreemux.ContextMux
	mid []Middleware
}

func (a *App) NotFoundHandler(h http.HandlerFunc) {
	a.mux.NotFoundHandler = h
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(mid ...Middleware) *App {
	apiMux := httptreemux.NewContextMux()

	return &App{
		mux: apiMux,
		mid: mid,
	}
}

// ServeHTTP implements the http. Handler interface. It's the entry point for
// all http traffic and allows the opentelemetry mux to run first to handle
// tracing. The opentelemetry mux then calls the application mux to handle
// application traffic.
func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}

// Handle sets a handler function for a given HTTP method and path pair
// to the application server mux.
func (a *App) Handle(method string, path string, handler Handler, mw ...Middleware) {

	// First wrap handler specific middleware around this handler.
	handler = wrapMiddleware(mw, handler)

	// Then wrap the application wide middlware around that
	handler = wrapMiddleware(a.mid, handler)

	// The function to execute for each request.
	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// if the error hasn't been handled by this point, it has propogated all the way out of our
		// call chain. If we're catching the error with a dedicated error middleware we should not
		// get to this point.
		if err := handler(ctx, w, r); err != nil {
			panic(fmt.Sprintf("uncaught error: %s", err))
		}
	}

	a.mux.Handle(method, path, h)
}

// HandleSocket is the same as Handle, except it doesn't add the gzip middleware. Gzip middleware
// does not implement http.Hijacker
func (a *App) HandleSocket(method string, path string, handler Handler, mw ...Middleware) {

	// First wrap handler specific middleware around this handler.
	handler = wrapMiddleware(mw, handler)

	// Then wrap the application wide middlware around that
	handler = wrapMiddleware(a.mid, handler)

	// The function to execute for each request.
	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// call the handler (with gzip if necessary)

		err := handler(ctx, w, r)

		// if the error hasn't been handled by this point, it has propogated all the way out of our
		// call chain. If we're catching the error with a dedicated error middleware we should not
		// get to this point.
		if err != nil {
			panic(fmt.Sprintf("uncaught error: %s", err))
		}
	}

	a.mux.Handle(method, path, h)
}
