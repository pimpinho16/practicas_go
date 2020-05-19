package main 

import ("net/http"
		"log")


func main(){
	/* es una forma básica de levantar un servidor 
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w,"Hola mundo desde mi servidor web con Go")
	})
	fmt.Println("Servidor corriendo desde puerto 8080")
	server := http.ListenAndServe(":8080",nil)
	*/
/*
	router := mux.NewRouter().StrictSlash(true) //todo buscar lo que hace la propiedad strictslash
	router.HandleFunc("/",Index)
	router.HandleFunc("/peliculas",MovieList)
	router.HandleFunc("/pelicula/{id}",MovieShow)
*/
	router := NewRouter()
	
	server := http.ListenAndServe(":8080",router)
	log.Fatal(server)


}

