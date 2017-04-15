package main

import (
	"counter/entriesofgo"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	log.Println(err, string(bytes))
	concurrency := 5
	sem := make(chan bool, concurrency)
	waiter := &sync.WaitGroup{}
	bytetostring := strings.TrimSpace(string(bytes))
	urls := strings.Split(bytetostring, "\n")
	infochan := make(chan entriesofgo.ResOfEntries, len(urls))
	for _, url := range urls {
		sem <- true
		waiter.Add(1)
		go entriesofgo.ScanForGo(url, infochan, sem, func() { waiter.Done() })

	}

	waiter.Wait()
	result := ""
	totalentries := 0

	close(infochan)
	for k := range infochan {
		result += k.Msg + "\n"
		totalentries += k.Cnt
	}
	result += fmt.Printf("Entries: ", totalentries)
	fmt.Println(result)

}
