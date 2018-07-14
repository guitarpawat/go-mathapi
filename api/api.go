package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/guitarpawat/go-mathapi/evaluate"
)

// APIHandler handles the http requst
func APIHandler(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.URL.Path, "/")
	result, err := evaluate.Evaluate(token[1:])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, result)
}
