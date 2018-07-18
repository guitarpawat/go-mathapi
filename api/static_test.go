package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func Test_IndexHandler_OK_DoctypeHTML(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	IndexHandler(res, req)
	resp := res.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code: %d but get: %d", http.StatusOK, resp.StatusCode)
		return
	}

	actual, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("unexpected error when reading body: %s", err)
		return
	}

	expected := "<!DOCTYPE html>"
	if !strings.ContainsAny(string(actual), expected) {
		t.Errorf("result does not contains: %s only get:\n%s", expected, string(actual))
	}
}

func Test_IndexHandler_NotFound(t *testing.T) {
	err := os.Rename("../README.md", "../README.tmp")
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	defer os.Rename("../README.tmp", "../README.md")

	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	IndexHandler(res, req)
	resp := res.Result()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("expected status code: %d but get: %d", http.StatusNotFound, resp.StatusCode)
		return
	}

}
