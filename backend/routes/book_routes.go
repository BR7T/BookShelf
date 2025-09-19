package routes

import (
	"fmt"
	"net/http"

	functions "github.com/BR7T/BookShelf/functions/book"
	"github.com/gorilla/mux"
)

func setupBookRoutes(router *mux.Router){
	userRouter := router.PathPrefix("/book").Subrouter()

	userRouter.HandleFunc("" , getBookByISBNhandler).Methods("GET")
	

}

func getBookByISBNhandler(w http.ResponseWriter , r *http.Request){

	err := functions.GetBookByISBN("9781648337178" , w , r)
	if err != nil{
		fmt.Fprint(w , err)
	}

}

