package main

import (
	"net/http"
	"log"

)

func SayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye,bye this is version 2" + r.URL.String()))
}
func main() {
	//使用mux
	mux := http.NewServeMux()

	mux.Handle("/",&myHandler{})
	mux.HandleFunc("/bye",SayBye)

	err := http.ListenAndServe(":8000",mux)
	if err != nil {
		log.Panic(err)
	}
}
type myHandler struct {}

func (_ *myHandler)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("Hello ,this is version 2 " + r.URL.String()))
}