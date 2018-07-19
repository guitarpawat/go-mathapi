package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/guitarpawat/go-mathapi/evaluate"
	"github.com/guitarpawat/go-mathapi/preprocess"
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

	token = preprocess.ConstantToNumber(token)

	if token[1] == "mem" {
		MemHandler(w, token[2:]...)
		return
	}

	if token[1] == "del" {
		if len(token) == 3 {
			DelHandler(w, token[2])
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "expected format: /del/name")
		}

		return
	}

	result, err := evaluate.Evaluate(token[1:])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, result)
}

// MemHandler handles for path /mem/...
func MemHandler(w http.ResponseWriter, infix ...string) {
	errMsg := "expected format: /mem/name/number or /mem/name/infix"
	if len(infix) < 2 {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, errMsg)
		return
	}

	err := preprocess.AddToMem(infix[0], infix[1:]...)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}

	value, _ := preprocess.GetFromMem(infix[0])

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Remember %s as %s", infix[0], value)
}

// DelHandler handles for path /mem/...
func DelHandler(w http.ResponseWriter, key string) {
	value, ok := preprocess.RemoveFromMem(key)
	if !ok {
		w.WriteHeader(http.StatusNoContent)
		fmt.Fprintf(w, "cannot find key: %s", key)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "successfully delete key %s with value %s", key, value)
}
