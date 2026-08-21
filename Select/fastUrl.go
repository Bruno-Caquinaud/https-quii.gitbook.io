package selects

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(firstUrl, secondUrl string, timeout time.Duration) (string, error) {

	select {
	case <-ping(firstUrl):
		return firstUrl, nil
	case <-ping(secondUrl):
		return secondUrl, nil
	case <-time.After(timeout * time.Second):
		return "", fmt.Errorf("Timeout Error")
	}
}

func ping(url string) chan struct{} {
	chanel := make(chan struct{})
	go func() {
		http.Get(url)
		close(chanel)
	}()

	return chanel
}

func measureResponseTime(url string) time.Duration {
	startFirst := time.Now()
	http.Get(url)
	return time.Since(startFirst)
}
