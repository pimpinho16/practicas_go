package main

import ("fmt"
		"os"
		"io/ioutil")	


func main(){
	nuevoTexto := os.Args[1]+"\n" //vendra un texto y se le agrega un salto de linea


	//fichero,errorDeFichero := ioutil.ReadFile("lenguajes.txt") para abrir el archivo solamente para su lectura
	fichero,errorDeFichero := os.OpenFile("lenguajes.txt",os.O_APPEND,0777) //le decimos al fichero que vamos a abrirlo para escriir en él, con todos los permisos (0777)

	showError(errorDeFichero)

	escribir,err := fichero.WriteString(nuevoTexto) //se obtiene el resultado de la operación WriteString
	fmt.Println(string(escribir))
	showError(err)

	fichero.Close()

	showFile,_ := ioutil.ReadFile("lenguajes.txt")

	fmt.Println(string(showFile))


}


func showError(e error){
	if(e != nil){
		panic(e) // corte la ejecución del programa y muestra el error
	}else {
		fmt.Println("NO hay error")
	}
}