package main

import (
	"net/http"
	"io"
)

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"foo lalllalal")
}
func bar(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"bar sfjoafiajfoiasf")
}
func main() {
	http.HandleFunc("/",foo)
	http.HandleFunc("/dog",bar)
	err := http.ListenAndServe(":8000",nil)
	if err != nil {
		panic(err)
	}
}