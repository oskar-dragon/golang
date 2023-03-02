package sel

import (
	"errors"
	"net/http"
	"time"
)

var ErrTimeout = errors.New("Timeout")

func Racer(url1, url2 string) (string, error) {


	select {
	case <- ping(url1):
		return url1, nil
	case <- ping(url2):
		return url2, nil
	case <- time.After(10 * time.Second):
		return "", ErrTimeout
	}
}

func ping(url string) chan string {
	ch := make(chan string)

	go func(u string) {
		http.Get(u)
		close(ch)
	}(url)

	return ch
}