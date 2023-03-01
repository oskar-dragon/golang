package concurrency

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string)bool {
	return url != "dupa.com"
}

func TestCheckWebsite(t *testing.T) {
	// create urls
	urls := []string{
		"http://google.com",
		"http://onet.pl",
		"dupa.com",
	}

	// create result
	want := map[string]bool{
		"http://google.com" : true,
		"http://onet.pl": true,
		"dupa.com": false,
	}
	// add test

	got := CheckWebsites(mockWebsiteChecker, urls)

	if !reflect.DeepEqual(want, got) {
		fmt.Printf("got %v, want %v",got, want)
	}
}

func slowedWebsiteChecker(_ string) bool {
	time.Sleep( 60 * time.Millisecond)

	return true
}

func BenchmarkCheckWebsite(b *testing.B) {
	urls := make([]string, 100)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowedWebsiteChecker, urls)
	}
}