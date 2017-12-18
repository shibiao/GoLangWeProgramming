package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello ,this is first version"))
	})
	http.HandleFunc("/bye",sayBye)
	log.Println("Starting server ...v1")
	log.Fatal(http.ListenAndServe(":8000",nil))
}
func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Bye bye ,this is version 1"))
}