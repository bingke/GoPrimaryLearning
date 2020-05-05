// Fetchall fetches same URL in parallel and compare their times and sizes.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main1_10() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
		go fetch(url, ch) //repeat
	}
	for range os.Args[1:] { //Avoid the main function exiting early
		fmt.Println(<-ch) //when the asynchronous execution of the goroutine has not been completed
	}
	fmt.Printf("%.2fs eplased", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	if strings.HasPrefix(url, "http://") != true {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	_, err = io.Copy(os.Stdout, resp.Body) //modify "ioutil.Discard" to "os.Stdout"
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("%s,%v", url, err)
		return
	}
	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("contents:%v, time:%.fs, url:%s", os.Stdout, secs, url) //put "os.Stdout" in channel

	//when I add os.Stdout,it can print all the contents in the web.
	//but this also bring a new problem,
	//it print context:&{0xc0000c4060}, time:2s, url:http://gopl.io ,why contents:&{0xc0000c4060}?

}
