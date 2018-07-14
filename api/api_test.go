package api

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_APIHandler_OK(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/5/mod/3", nil)
	APIHandler(res, req)
	resp := res.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code: %d but get: %d", http.StatusOK, resp.StatusCode)
		return
	}

	actual, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("unexpected error when reading body: %s", err)
	}

	expected := "2.0000"
	if strings.TrimSpace(string(actual)) != expected {
		t.Errorf("expected result:|%s| but get:|%s|", expected, actual)
	}
}

func Test_APIHandler_Error(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/5/fun/3", nil)
	APIHandler(res, req)
	resp := res.Result()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status code: %d but get: %d", http.StatusInternalServerError, resp.StatusCode)
	}
}
