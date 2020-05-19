package main

import ("fmt"
		"os"
		"strconv")


func main (){
	numero,_ := strconv.Atoi(os.Args[1])

	if (numero % 2 == 0 ){
		fmt.Println("El número es par")
	}else if (numero % 2 == 1){
		fmt.Println("El número es impar")
	}else {
		fmt.Println("Error")
	}

}			