package main 

import "fmt"

func main(){
	arrays()

} 


func arrays(){
	var peliculas [3]string;
	peliculas[0] = "green mile"
	peliculas[1] = "die hard"
	peliculas[2] = "Rocky"

	fmt.Println(peliculas)
	fmt.Println(peliculas[2])


	peliculas2 := [3]string {
		"1917",
		"2 popes",
		"cage"} 
	fmt.Println(peliculas2)
	fmt.Println(peliculas2[1])	
	
}