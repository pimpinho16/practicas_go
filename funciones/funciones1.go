package main 

import "fmt"

func main(){
	imprimirTexto("Hola Mundo, esta es una calculadora")
	calculadora(10,20,"+")
	calculadora(10,20,"-")
	calculadora(10,20,"*")
	calculadora(10,20,"/")
}


func imprimirTexto(text string) {
	fmt.Println(text)
}

func calculadora(op1 float32, op2 float32, operador string ) {
	var resultado float32
	if(operador == "+"){
		resultado = op1 + op2
	}else if(operador == "-"){
		resultado = op1 - op2
	}else if (operador == "*"){
		resultado = op1 * op2

	}else if (operador == "/"){
		resultado = op1/op2
	}else {
		resultado =0;
	}
	imprimirTexto("El resultado es " )
	fmt.Println(resultado)
}

