package routes

import (
	"fmt"
	_ "io"
	"net/http"

	functions "github.com/BR7T/BookShelf/functions/user"
	"github.com/BR7T/BookShelf/models"
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
	users , err := functions.GetAllUsers()
	if err != nil{
		fmt.Fprint(w, err)
	}
	fmt.Printf("\n Usuários encontrados: \n %+v" , users)
	
}

func addUserHandle(w http.ResponseWriter , r *http.Request){
	var user models.UserRegister
	err := utils.ParseJsonBody(w,r, &user)
	if err != nil{
		fmt.Fprint(w,err)
	}
	service.RegisterUser(user)
}

func loginUserHandle(w http.ResponseWriter , r *http.Request){
	var login models.UserLogin
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