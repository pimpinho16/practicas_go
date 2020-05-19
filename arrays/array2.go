package main 

import "fmt"

func main (){

	arrays()
}

func arrays(){
	var peliculas [3][2]string;
	peliculas[0][0] = "green mile"
	peliculas[0][1] = "pursuit of happiness"
	
	peliculas[1][0] = "Rocky"
	peliculas[1][1] = "terminator"

	peliculas[2][0] = "la lista de shindler"
	peliculas[2][1] = "el niño con pijama de rayas"

	for _, pelicula := range peliculas[0]{
		fmt.Println(pelicula)
	} 
}


