package main 

import ("net/http"
		"github.com/gorilla/mux"
		"fmt"
		"encoding/json"
		"log"
		"gopkg.in/mgo.v2/bson")
	//	"log"
	//	"gopkg.in/mgo.v2")
		//

/*
var moviesGlobal = Movies{
		Movie{"Sin Limite",2013,"sepa putas"},
		Movie{"Amores Perros",2000,"Alejandro Gonzalez"}}

*/


//se invoca la conexión a mongo db para esa base y esa coleccion 
var collection = getSession().DB("curso_go").C("movies")	

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Servidor corriendo desde puerto 8080")
}

func responseMethod(w http.ResponseWriter, httpStatus int,results Movie){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

func responseMethods(w http.ResponseWriter, httpStatus int,results []Movie){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}


func MovieList(w http.ResponseWriter, r *http.Request){
	var results []Movie //se define un array de Movie en donde se almacenará el resultado de la consulta en mongodb
	err := collection.Find(nil).Sort("-_id").All(&results) //en find como se desea obtener todos los resultdos se manda nil es decir, no queremos filtrar 

	if(err!= nil){
		log.Fatal(err)
	}else{
		fmt.Println("Resultados",results)
	}
	//fmt.Fprintf(w,"Esta es la página del listado de películas")
	
	//json.NewEncoder(w).Encode(moviesGlobal) //mando en NewEncoder el w del response writer y en el encode el objeto que quiero transformar a json

	responseMethods(w,200,results) //se simplifica el código para mandar el resultado 
}

/*
func MovieShow(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	movie_id := params["id"] //array asociativo se consulta através de su llave, en este caso es id
	fmt.Fprintf(w,"Has cargado la película %s",movie_id)	
}
*/


func MovieShow(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	movie_id := params["id"] //array asociativo se consulta através de su llave, en este caso es id
	if !bson.IsObjectIdHex(movie_id){ //se verifica si el id es un object id hexadecimal
		w.WriteHeader(404)
		return
	}
	
	oid := bson.ObjectIdHex(movie_id) // ObjectIdHex returns an ObjectId from the provided hex representation.
	results := Movie{}

	err := collection.FindId(oid).One(&results)//se obtiene el valor de la coleccion a apartir del object id

	if(err != nil){
		w.WriteHeader(404)
		return
	}
	responseMethod(w,200,results) //se simplifica el código para mandar el resultado 

	

}

func MovieAdd(w http.ResponseWriter, r *http.Request){
	decoder  := json.NewDecoder(r.Body)

	var movie_data Movie
	err := decoder.Decode(&movie_data) // aquí se le asigna a movie data lo que viene en el body, por eso se le manda la referencia del campo con &

	if(err != nil){
		panic(err)
	}

	defer r.Body.Close() // esto limpia las lecturas del body

	//log.Println(movie_data) //imprimo lo que viene en el post
	//moviesGlobal = append(moviesGlobal, movie_data) // agrego al array global el nuevo elemento

	//session := getSession()
	err = collection.Insert(movie_data)

	if(err != nil){ // si err es diferente a nil es que no ha podido realizarse la inserción
		w.WriteHeader(500) //se devuelve un código http 500
		return //se retorna automaticamente para que la función ya n continue
	}
	//devolver un json
	
	
	responseMethod(w,200,movie_data) //se simplifica el código para mandar el resultado 
}


func MovieUpdate(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	movie_id := params["id"] //array asociativo se consulta através de su llave, en este caso es id
	
	if !bson.IsObjectIdHex(movie_id){ //se verifica si el id es un object id hexadecimal
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(movie_id) // ObjectIdHex returns an ObjectId from the provided hex representation.
	
	decoder := json.NewDecoder(r.Body)
	
	var movie_data Movie

	err := decoder.Decode(&movie_data)

	if (err != nil){
		panic(err)
		w.WriteHeader(500)
		return
	}

	defer r.Body.Close()


	document := bson.M{"_id":oid} // se obtiene el objeto en base al oid en formato bson

	change := bson.M{"$set":movie_data} //creamos un bson con la info que viene en el request de la peticion

	err = collection.Update(document,change) //para actualizar se manda el document con el objeto actual y el formato bson del objeto con las modificacione

	if (err != nil){
		panic(err)
		w.WriteHeader(404)
		return
	}

	responseMethod(w,200,movie_data) //se simplifica el código para mandar el resultado 

}


func MovieDelete(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	movie_id := params["id"] //array asociativo se consulta através de su llave, en este caso es id
	if !bson.IsObjectIdHex(movie_id){ //se verifica si el id es un object id hexadecimal
		w.WriteHeader(404)
		return
	}
	
	oid := bson.ObjectIdHex(movie_id) // ObjectIdHex returns an ObjectId from the provided hex representation.

	err := collection.RemoveId(oid)//se obtiene el valor de la coleccion a apartir del object id

	if(err != nil){
		w.WriteHeader(404)
		return
	}
	
	//results := Message{"sucess","La película con id"+movie_id+" ha sido borrada correctamente" }

	message := new(Message)

	message.SetStatus("sucess")
	message.SetMessage("La película con id"+movie_id+" ha sido borrada correctamente") 
	results := message
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
	//responseMethod(w,200,results) //se simplifica el código para mandar el resultado 

	

}