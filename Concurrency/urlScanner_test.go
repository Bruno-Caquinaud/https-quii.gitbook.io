package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {

	if url == "waat://furhurterwe.geds" {
		return false
	}

	return true
}

func stubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func TestCheckWebsite(t *testing.T) {
	urls := []string{
		"http://google.com",
		"http://blog.gypsydave5.com",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":          true,
		"http://blog.gypsydave5.com": true,
		"waat://furhurterwe.geds":    false,
	}

	got := CheckWebsite(mockWebsiteChecker, urls)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got : %v != want : %v", got, want)
	}
}

func BenchmarkWebsiteChecker(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < 100; i++ {
		urls[i] = "a url"
	}

	for b.Loop() {
		CheckWebsite(stubWebsiteChecker, urls)
	}
}
