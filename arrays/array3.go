package main

import "fmt"


func main(){
	slice()
}

/* slice es un array sin corchete, es decir es un array dinámico */  
func slice(){
	peliculas := []string{
		"la verdad duele",
		"el crimen del padre amaro",
		"y tu mama tambien"}
	fmt.Println(peliculas)
		
	peliculas = append(peliculas,"amores perros")	//agregar elementos dinamicamente
	fmt.Println(peliculas)

	fmt.Println(len(peliculas)) //obtiene la longitud de array o slice 

	fmt.Println(peliculas[0:2]) //	quiero los elementos del 0 al 2
}