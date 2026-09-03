package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {

	channel := make(chan string, 1)

	go func() {
		var result string

		for _, c := range s.response {
			select {
			case <-ctx.Done():
				return
			default:

				result += string(c)
			}
		}
		channel <- result
	}()

	time.Sleep(10 * time.Millisecond)
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-channel:
		return res, nil
	}
}

func TestServer(t *testing.T) {
	t.Run("Test Server 1 : Running valid handling of HTTP request func", func(t *testing.T) {
		data := "Hello, World"
		spystore := &SpyStore{data}
		srv := Server(spystore)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		srv.ServeHTTP(response, request)

		assertString(t, response.Body.String(), data)
	})

	t.Run("Test Server 2 : Running a handler HTTP Request calling cancellation context", func(t *testing.T) {
		data := "Hello, World"
		spystore := &SpyStore{data}
		srv := Server(spystore)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		cancelingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancelingCtx)

		srv.ServeHTTP(response, request)

		assertString(t, response.Body.String(), "")
	})

}

func assertString(t *testing.T, want, expected string) {
	if want != expected {
		t.Errorf("Want : %s != %s expected", want, expected)
	}
}
