package routes

import (
	"fmt"
	"io"
	"net/http"

	"github.com/BR7T/BookShelf/service"
	"github.com/gorilla/mux"
)

func setupUserRoutes(router *mux.Router){
	userRouter := router.PathPrefix("/users").Subrouter()

	userRouter.HandleFunc("" , getUsersHandle).Methods("GET")
	userRouter.HandleFunc("" , addUserHandle).Methods("POST")

}

func getUsersHandle(w http.ResponseWriter , r *http.Request){
	fmt.Print("Buscando Usuários")
	users , err := service.GetAllUsers()
	if err != nil{
		fmt.Fprint(w, err)
	}
	fmt.Printf("\n Usuários encontrados: \n %+v" , users)
	
}

func addUserHandle(w http.ResponseWriter , r *http.Request){
	body , _ := io.ReadAll(r.Body)
	fmt.Print(body)
}