package main

type Message struct{
	Status string `json:"status"`
	Message string `json:"message"`
}

func (this *Message) SetStatus(data string){ //se deben recibir como referencia para manipular un objeto ya existente
	this.Status = data
}


func (this *Message) SetMessage(data string){
	this.Message = data
}