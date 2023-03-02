package sel

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		server1 := makeDelayedServer(1 * time.Millisecond)
		server2 := makeDelayedServer(10 * time.Millisecond)

		defer server1.Close()
		defer server2.Close()

		url1 := server1.URL
		url2 := server2.URL

		got,_ := Racer(url1, url2)

		want := url1
		
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("times out after 10 seconds", func (t *testing.T) {
		server1 := makeDelayedServer(11 * time.Second)
		server2 := makeDelayedServer(12 * time.Second)

		defer server1.Close()
		defer server2.Close()

		url1 := server1.URL
		url2 := server2.URL

		_,err := Racer(url1, url2)

		want := ErrTimeout

		if err != want {
			t.Errorf("got error %s but want error %s", err, want)
		}

	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}