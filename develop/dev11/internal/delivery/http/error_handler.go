package http

import (
	"fmt"
	"net/http"
)

func httpMethodErrorCheck(w http.ResponseWriter, expectedMethod, gotMethod string) bool {
	if gotMethod != expectedMethod {

		http.Error(
			w,
			fmt.Sprintf("expected: %s, got instead: %s", expectedMethod, gotMethod),
			http.StatusBadRequest)
	}
	return gotMethod != expectedMethod
}
