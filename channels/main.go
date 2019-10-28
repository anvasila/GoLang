package main

import (
	"fmt"
	"net/http"
	"time"
)

func main(){

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://interplast.gr",
	}

	// Make Channel
	c := make(chan string)

	for _,link := range links{
		// Create a Go Routine to run code parallel
		go checkLink(link, c)
	}

	// W8 for all urls to get result and then end
	//for i:=0; i<len(links);i++{
	//	fmt.Println( <- c )
	//}

	// Do it forever
	//for {
	//	go checkLink(<-c,c)
	//}

	// Wait for the channel to return some value
	// Then assign it to l and run the body of the for loop
	for l := range c{
		//go checkLink(l,c)
		go func(link string) {
			time.Sleep(5 * time.Second)
			go checkLink(link,c)
		}(l)
	}

	fmt.Println("Done!")
}

func checkLink(link string, c chan string){
	_,err := http.Get(link)

	if err != nil{
		fmt.Println(link,"might be down!")
		c <- link
	}

	fmt.Println(link,"is up!")
	c <- link
}