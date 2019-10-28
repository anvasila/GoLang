package main

import "fmt"

// Because both englishBot and spanishBot implement
// function getGreeting automatically make them be type bot
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main(){
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hi There!!!"
}

func (spanishBot) getGreeting() string {
	return "Hola!!!"
}