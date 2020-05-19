package main

import "fmt"

type Gorra struct {
	marca string
	color string
	precio float32
	plana bool
}

func main(){
	var gorra_negra = Gorra{"adidas", "Negra",12.5, false}

	gorra_negra.marca = "nike"
	fmt.Println(gorra_negra)
}