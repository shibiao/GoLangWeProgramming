package main

import (
	"net/http"
	"log"
	"time"
	"os"
	"os/signal"
)

func SayBye(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Bye,bye this is version 3" + r.URL.String()))
}
func main() {
	//添加server
	server := http.Server{
		Addr:":8000",
		WriteTimeout:2*time.Second,
		ReadTimeout:2*time.Second,
	}
	quit := make(chan os.Signal)
	signal.Notify(quit,os.Interrupt)
	//使用mux
	mux := http.NewServeMux()

	mux.Handle("/",&myHandler{})
	mux.HandleFunc("/bye",SayBye)
	server.Handler = mux

	go func() {
		<- quit
		if err := server.Close(); err != nil {
			log.Fatal("Close server:",err)
		}
	}()
	log.Println("Starting server .... v3")
	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Print("Server closed under request")
		}else{
			log.Print("Server closed under unexpected")
		}
	}
	log.Println("Sever exit")
}
type myHandler struct {}

func (_ *myHandler)ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	w.Write([]byte("Hello ,this is version 3 " + r.URL.String()))
}