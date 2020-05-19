package main //el main del proyecto inicia con ese nombre de package

import ("fmt"
		"time"
		)
		 //así se llaman paquetes en Go

func main(){
	fmt.Println("Hola mundo desde Go") //para escribir cadenas se deben colocar en comillas dobles
	time.Sleep(time.Second * 5)
} 