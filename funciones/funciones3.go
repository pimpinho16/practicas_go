package main

import "fmt"

func main(){
	fmt.Println(calcularTotal(10,2.5))	
}

func calcularTotal(unidades float32, precio float32)(string,float32){
	total := func() (float32){
		return unidades*precio
	}
	return "El resultado es ", total()

}