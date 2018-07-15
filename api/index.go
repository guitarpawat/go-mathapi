package api

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang-commonmark/markdown"
)

const (
	header = `<!DOCTYPE html>
<head><title>Go-MathAPI</title></head>
<body>
`
	footer = `
</body>
</html>`
)

// IndexHandler handles http request of root page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	md, err := ioutil.ReadFile("README.md")
	if err != nil {
		// For testing env
		md, err = ioutil.ReadFile("../README.md")
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "404 not found")
			return
		}
	}

	output := markdown.New(markdown.HTML(true))
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s%s%s", header, output.RenderToString(md), footer)
}
