package httpx

import (
	"fmt"
	"net/http"

	"github.com/dimfeld/httptreemux"
)

// Param returns the web call parameters from the request.
func PathParam(r *http.Request, key string) string {
	vars := httptreemux.ContextParams(r.Context())
	param, ok := vars[key]
	if !ok {
		panic(fmt.Sprintf("missing path param: path parameter %q does not exist on route", key))
	}

	return param
}

// QueryParam is a convenient method for getting a query param in the URL
func QueryParam(r *http.Request, key string) (string, bool) {
	ok := r.URL.Query().Has(key)
	param := r.URL.Query().Get(key)

	return param, ok
}
