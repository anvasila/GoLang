package main

import "fmt"

type person struct {
	firstName string
	lastName string
	contactInfo
}

type contactInfo struct {
	email string
	zipCode int
}

func main(){
	tasos := person{
		firstName: "Tasos",
		lastName:  "Vasilakakis",
		contactInfo:   contactInfo{
			email: "anvasil@hotmail.com",
			zipCode: 69100,
		},
	}
	tasos.updateName("Anastasios")
	tasos.print()
}

func (p *person) updateName(newFirstName string){
	p.firstName = newFirstName
}

func (p person) print(){
	fmt.Printf("%+v",p)
}