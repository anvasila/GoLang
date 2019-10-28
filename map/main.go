package main

import "fmt"

func main(){

	// Option #1
	colors := map[string]string{
		"red": "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)

	// Option #2
	//var colors map[string]string

	// Option #3
	//colors := make( map[string]string )
	//colors["white"] = "#ffffff"
	//delete(colors,"white")

	// fmt.Println(colors)
}

func printMap( c map[string]string ){
	for color, hex := range c{
		fmt.Println("Color "+color+" is "+hex)
	}
}