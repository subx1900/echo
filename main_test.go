package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestEcho(t *testing.T) {
	t.Run("Tests echo3's empty body", func(t *testing.T) {

		request, _ := http.NewRequest(http.MethodGet, "/", bytes.NewBufferString(url.Values{}.Encode()))
		response := httptest.NewRecorder()

		echo(response, request)

		got := response.Body.String()
		gotStruct := &Response{}

		_ = json.Unmarshal([]byte(got), gotStruct)

		if string(gotStruct.Body) != "" {
			t.Errorf("got %q, want %q", got, "''")
		}
	})
}