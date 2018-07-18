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
		t.Errorf("expected result: %s but get: %s", expected, actual)
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

func Test_APIHandler_ExtraSlash_SeeOther(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/5/mod/3/", nil)
	APIHandler(res, req)
	resp := res.Result()

	if resp.StatusCode != http.StatusSeeOther {
		t.Errorf("expected status code: %d but get: %d", http.StatusSeeOther, resp.StatusCode)
		return
	}
}

func Test_APIHandler_Mem(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/mem/answer/42", nil)
	APIHandler(res, req)
	resp := res.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status code: %d but get: %d", http.StatusOK, resp.StatusCode)
		return
	}
}

func Test_APIHandler_Slash_ToIndexHandler(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	APIHandler(res, req)
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

func Test_APIHandler_Call_ConstantToNumber(t *testing.T) {
	res := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/pi", nil)
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

	expected := "3.141593"
	if strings.TrimSpace(string(actual)) != expected {
		t.Errorf("expected result: %s but get: %s", expected, actual)
	}
}

func Test_MemHandler(t *testing.T) {
	res := httptest.NewRecorder()
	expected := "Remember ant as 1"

	MemHandler(res, "ant", "1")

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

	if strings.TrimSpace(string(actual)) != expected {
		t.Errorf("result does not contains: %s only get:\n%s", expected, string(actual))
	}
}

func Test_MemHandler_Invalid(t *testing.T) {
	res := httptest.NewRecorder()

	MemHandler(res, "ant")

	resp := res.Result()
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status code: %d but get: %d", http.StatusInternalServerError, resp.StatusCode)
	}
}

func Test_MemHandler_Reject(t *testing.T) {
	res := httptest.NewRecorder()

	MemHandler(res, "12", "34")

	resp := res.Result()
	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("expected status code: %d but get: %d", http.StatusInternalServerError, resp.StatusCode)
	}
}
