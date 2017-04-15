package entriesofgo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ResOfEntries struct {
	Msg string
	Cnt int
}

func getHtmlbyUrl(url string) (htmlData string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error occurred while getting the content", err)
		return
	}
	defer resp.Body.Close()
	// Reads html as a slice  of bytes.
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error occurred while reading html", err)
		return
	}
	htmlData = string(bytes)
	return

}

func ScanForGo(url string, chanres chan ResOfEntries, sem chan bool, done func()) {
	if done != nil {
		<-sem
		defer done()
	}
	datahtml, err := getHtmlbyUrl(url)
	if err != nil {
		fmt.Println("Scan failed", err)
	}
	cnt := strings.Count(datahtml, "Go")
	chanMsg := fmt.Sprintf("Count for  : ", url, cnt)
	chanres <- ResOfEntries{Msg: chanMsg, Cnt: cnt}
}
