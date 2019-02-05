package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://facebook.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	for l := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "Error Uploading Link")
		c <- link
		return
	}

	fmt.Println(link, " is working")
	c <- link

}
