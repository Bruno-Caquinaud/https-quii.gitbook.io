package selects

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("Test Racer 1 : Check fastest GET Request URL is return", func(t *testing.T) {
		slowServer := makeMockServer(20)
		fastServer := makeMockServer(0)
		defer slowServer.Close()
		defer fastServer.Close()

		got, _ := Racer(slowServer.URL, fastServer.URL, 10)
		want := fastServer.URL

		if got != want {
			t.Errorf("got %q != want %q", got, want)
		}
	})

	t.Run("Test Racer 2 : Check Raised Error when GET URL take more than 10 seconds", func(t *testing.T) {
		slowServer := makeMockServer(11000)
		fastServer := makeMockServer(11000)
		defer slowServer.Close()
		defer fastServer.Close()

		_, err := Racer(slowServer.URL, fastServer.URL, 1)

		if err == nil {
			t.Errorf("Error expected but no one has been raised")
		}
	})

}

func makeMockServer(timesleep time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(timesleep * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
}
