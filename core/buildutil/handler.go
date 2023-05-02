package buildutil

import (
	"encoding/json"
	"net/http"
)

func VersionHandler() http.HandlerFunc {

	// Convert the response value to JSON at startup time, then we only
	// have to marshal the data once
	jsonData, err := json.Marshal(InfoEmbed)
	if err != nil {
		panic(err.Error())
	}

	h := func(w http.ResponseWriter, r *http.Request) {

		// Set the content type and headers once we know marshaling has succeeded.
		w.Header().Set("Content-Type", "application/json")

		// Write the status code to the response.
		w.WriteHeader(http.StatusOK)

		// Write response data to response body.
		_, err = w.Write(jsonData)
		if err != nil {
			panic(err.Error())
		}

	}

	return h
}
