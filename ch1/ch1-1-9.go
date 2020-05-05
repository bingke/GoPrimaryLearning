// Modify fetch to print out the status code of the HTTP protocol
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main1_9() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://") != true {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b := resp.Status //The status code can be obtained from the resp.Status variable
		resp.Body.Close()

		fmt.Printf("%s", b)
	}
}
