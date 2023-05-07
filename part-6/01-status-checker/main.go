package main

import (
	"fmt"
	"net/http"
)

func main() {
	websites := []string{"http://facebook.com", "http://amazon.co.uk", "http://youtube.com"}

	c := make(chan string)

	for _, link := range websites {
		go checkWebsite(link, c)
	}

	for i := 0; i < len(websites); i++ {
		fmt.Println(<-c)
	}

}

func checkWebsite(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is down!")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is ok")
}
