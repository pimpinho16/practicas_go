package main 

import ("fmt"
		"os" // esta libreria sirve para interactuar con argumentos de consola
		"strconv") //


func main (){

	
	nombre := os.Args[1]
	edad,_ := strconv.Atoi(os.Args[2]) //Atoi->convierte string a entero
	
	fmt.Println("Bienvenido " + nombre+ ". Tienes " + strconv.Itoa(edad)+ " años" ) //convierte string a entero
	if (edad >=18 || edad ==17 )  && edad != 25 && edad != 99{
		fmt.Println ("Eres mayor de edad o tienes 17 años")
	} else if edad == 25 {
		fmt.Println("Tienes 25 años")

	}else if  edad == 99 {
			fmt.Println("Pronto moriras")
	}else {
		fmt.Println("Eres menor de edad")
	}

	
}