package web

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
