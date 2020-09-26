package test

import (
	"LearnGoGin/Controller"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowIndexPage(t *testing.T) {
	r := getRouter(true)

	r.GET("/", Controller.ShowIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {
		statusOK := w.Code == http.StatusOK

		p, err := ioutil.ReadAll(w.Body)
		pageOk := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOk
	})
}
