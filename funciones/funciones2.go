package main

import "fmt"

func main(){
	fmt.Println(devolverTexto("Alexander",99))
}

func devolverTexto(dato1 string, dato2 int) (dato3 string, dato4 int){
	dato3 = dato1
	dato4 = dato2
	return //go sabe que va a devolver los valores que agregamos en los segundos parentesis