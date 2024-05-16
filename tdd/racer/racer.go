package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	//durationA := measureResponseTime(a)
	//durationB := measureResponseTime(b)

	//if durationA < durationB {
	//	return a
	//}

	//return b

	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

//func measureResponseTime(url string) time.Duration {
//	start := time.Now()
//	http.Get(url)
//	return time.Since(start)
//}

// ping creates a chan struct{} and closes it once a GET request
// to the provied URL completes. Why chan struct{}? In our case,
// we don't actually care about the return value of the GET, only
// that it completed. chan struct{} is the smallest data type
// available, from a memory perspective, so we get no allocation
// versus a type like bool. Why allocate any memory if we don't
// care about the value?
func ping(url string) chan struct{} {
	ch := make(chan struct{})

	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
