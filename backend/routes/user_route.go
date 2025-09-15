package routes

import (
	"fmt"
	_ "io"
	"net/http"

	"github.com/BR7T/BookShelf/service"
	"github.com/BR7T/BookShelf/utils"
	"github.com/gorilla/mux"
)

func setupUserRoutes(router *mux.Router){
	userRouter := router.PathPrefix("/users").Subrouter()

	userRouter.HandleFunc("" , getUsersHandle).Methods("GET")
	userRouter.HandleFunc("" , addUserHandle).Methods("POST")
	userRouter.HandleFunc("/login" , loginUserHandle).Methods("POST")

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
	var user service.UserRegister
	err := utils.ParseJsonBody(w,r, &user)
	if err != nil{
		fmt.Fprint(w,err)
	}
	service.RegisterUser(user)
}

func loginUserHandle(w http.ResponseWriter , r *http.Request){
	var login service.UserLogin
	err := utils.ParseJsonBody(w,r, &login)
	if err != nil{
		fmt.Fprint(w,err)
	}

	valid , err := service.LoginUser(login)
	if err != nil{
		fmt.Print(err)
	}

	if valid{
		w.Write([]byte("Entrou"))
	}else{
		w.Write([]byte("Não Entrou"))
	}
}