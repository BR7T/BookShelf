package main

import(
	"github.com/BR7T/BookShelf/routes"
	"net/http"
	"log"
)

func main(){
	router := routes.InitRouter()
	
	port := ":8080"
    log.Printf("Servidor iniciado na porta %s", port)
    log.Fatal(http.ListenAndServe(port, router))
}