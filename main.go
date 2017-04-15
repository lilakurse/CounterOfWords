package main

import (
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
	scanWG := &sync.WaitGroup{}
	bytetostring := strings.TrimSpace(string(bytes))
	urls := strings.Split(bytetostring, "\n")
	//result_channel := make(chan ,len(urls))
	for _, url := range urls {
		sem <- true
		scanWG.Add(1)
		//go numofentries()

	}

	scanWG.Wait()
	result := ""
	totalentries := 0

	//close(result_channel)
	/* for k :=range result_channel{
		result +=
		totalentries +=
	}
	result += fmt.Printf("Entries: ",totalentries)
	fmt.Println(result)
	*/

}

