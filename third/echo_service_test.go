package isupov

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEcho(t *testing.T) {
	t.Run("POST test string", func(t *testing.T) {
		data := "TEST"
		request, _ := http.NewRequest(http.MethodPost, "/", strings.NewReader(data))

		response := httptest.NewRecorder()
		Echo(response, request)

		got := response.Body.String()
		if got != data {
			t.Errorf("got: %v, expected: %v", got, data)
		}
	})

	t.Run("Wrong HTTP method", func(t *testing.T) {
		data := "TEST"
		request, _ := http.NewRequest(http.MethodGet, "/", strings.NewReader(data))

		response := httptest.NewRecorder()
		Echo(response, request)

		if response.Code != http.StatusMethodNotAllowed {
			t.Error("only POST method is allowed")
		}
	})
}
