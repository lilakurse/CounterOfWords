// Package test the work of program
package entriesofgo

import (
	"sync"
	"testing"
)

func TestScanForGo(t *testing.T) {
	urls := []string{
		"https://golang.org",
		"https://www.tutorialspoint.com/go/",
		"https://tour.golang.org/welcome/1",
	}

	sem := make(chan bool, 3)
	infoChannel := make(chan ResOfEntries, 3)
	waiter := &sync.WaitGroup{}
	for _, url := range urls {
		sem <- true
		waiter.Add(1)
		go ResOfEntries{}(url, infoChannel, sem, waiter)
	}
	waiter.Wait()

	if len(infoChannel) != 3 {
		t.Error("Error occurred")
	}
	close(infoChannel)
	for res := range infoChannel {
		if res.Cnt == 0 {
			t.Error("Can not find the line GO")
		}
	}
}
