package main

import "fmt"

func main (){
	for i:=1 ; i < 16; i++{
		if( i % 2 ==0)	{
			fmt.Println("Numero par: ",i )
		}else {
			fmt.Println("Numero impar: ",i )
		}
	}
}