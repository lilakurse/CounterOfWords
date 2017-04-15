package entriesofgo

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getHtmlbyUrl(url string) {
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
	html := string(bytes)

}
