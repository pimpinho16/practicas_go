package main //todos los archivos con el paquete main son accesibles entre si

type Movie struct{
	Name string     `json:"nombre"` //de esta forma formateamos que el json muestre la variable en minuscula
	Year int        `json:"anio"`
	Director string `json:"director"`
}

type Movies []Movie