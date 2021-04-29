package server_test

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/kenhowardpdx/ecommerce-ratings/internal/server"
)

func TestHandleIndex(t *testing.T) {
	t.Run("serve version on index", func(t *testing.T) {
		want := "ecommerce-ratings 1.2.3\n"
		srv := server.Server{
			Version: "1.2.3",
		}
		routes := srv.Routes()
		w := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/", nil)
		routes.ServeHTTP(w, req)
		got, err := ioutil.ReadAll(w.Result().Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(got) != want {
			t.Errorf("got: %q, want: %q", got, want)
		}
	})
}
