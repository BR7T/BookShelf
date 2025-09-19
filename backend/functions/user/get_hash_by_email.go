package functions

import (
	"fmt"
	"github.com/BR7T/BookShelf/config"
	"github.com/BR7T/BookShelf/models"
)

func GetHashPassword(email string)(*models.UserPassword , error){
	var user models.UserPassword

	db := config.SetupDB()
	err := db.QueryRow("SELECT email , password , role FROM users WHERE email = $1" , email).Scan(&user.Email , &user.Password , &user.Role)
	if err != nil{
		return nil , fmt.Errorf("erro ao fazer a consulta sql: %v" , err)
	}

	return &user , nil
}