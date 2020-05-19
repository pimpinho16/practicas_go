package main

import "fmt"

func main(){
	pantalon("azul","punta de yuca","36 de cintura","sin bolsillos traseros","pierre cardin")
}

func pantalon (caracteristicas ...string){ //forma de go para recibir parámetros dinámicos a una función
	for _,caracteristica := range caracteristicas{ //for que simula el foreach donde se recorre un objeto o array
		fmt.Println(caracteristica)
	}
}	