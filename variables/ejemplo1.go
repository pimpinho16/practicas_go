package main

import "fmt"

func main() {
	var suma int = 8 + 9
	var nombre string = "Alexander"


	fmt.Println(suma)
	fmt.Println("Mi nombre es " + nombre)
	nombre = "Alberto"
	fmt.Println("Mi nombre es " + nombre)
	
	apellido := "Andrade" //Go infiere que el tipo de datos será string, no puedo darle valor diferente de un string

	fmt.Println(apellido)

	var otro float32 = 1.32
	siono  := false
	fmt.Println(otro)
	fmt.Println(siono)

	const year int = 2018 //si se intenta asignar en caliente un valor nuevo a year el compilador no dejará 
	fmt.Println(year)
}
