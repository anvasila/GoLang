package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type logWriter struct{}

func main(){
	resp, err := http.Get("http://google.com/")
	if err != nil {
		log.Fatal(err)
	}

	//bs := make([]byte,999999)
	//resp.Body.Read(bs)
	//fmt.Println( string(bs) )

	// Copy(dst Writer, src Reader)
	// So we pass as Reader the request body (witch is a Reader)
	// and as Write to os.Stdout(to copy it in terminal)
	//io.Copy(os.Stdout, resp.Body)

	// We create a struct with function Write (so it is implements Writer Interface)
	// This way we can use a struct for Writer
	lw := logWriter{}
	io.Copy(lw,resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:",len(bs))
	return len(bs),nil
}