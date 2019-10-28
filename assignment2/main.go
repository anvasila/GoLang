package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct{
	height float64
	base float32
}
type square struct{
	sideLength float64
}

func main(){
	tr := triangle{10,10}
	sq := square{10}

	printArea(tr)
	printArea(sq)
}

func printArea(s shape){
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64{
	return float64(0.5)*t.height*float64(t.base)
}

func (s square) getArea() float64{
	return s.sideLength * s.sideLength
}