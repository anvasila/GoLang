package main

import (
	//"fmt"
	"io"
	//"io/ioutil"
	"log"
	"os"
)

func main(){
	args := os.Args
	fileName := args[ len(args)-1 ]

	// With ReadFile
	//content,err := ioutil.ReadFile(fileName)
	//
	//if err != nil{
	//	fmt.Println("We got err",err)
	//}
	//
	//fmt.Println( string(content) )

	// With os.Open
	file, err := os.Open(fileName)

	if err != nil{
		log.Fatal(err)
	}

	io.Copy(os.Stdout,file)

}