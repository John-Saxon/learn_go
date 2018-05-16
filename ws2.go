package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string


type Struct struct{
	Greeting string
	Punct string
	Who string
}


func (str String) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request){
	fmt.Fprint(w, str)
}

func (stt Struct) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request){
	fmt.Fprint(w, fmt.Sprintf("%s,%s!",stt.Greeting,stt.Who) )
}

func main() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:4000", nil))
}

