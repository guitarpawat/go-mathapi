package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/guitarpawat/go-mathapi/evaluate"
)

// APIHandler handles the http requst
func APIHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		IndexHandler(w, r)
		return
	}

	if string(r.URL.Path[len(r.URL.Path)-1]) == "/" {
		http.Redirect(w, r, string(r.URL.Path[:len(r.URL.Path)-1]), http.StatusSeeOther)
		return
	}

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
