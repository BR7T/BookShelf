package functions

import (
	"database/sql"

	"github.com/BR7T/BookShelf/config"
)

func ExistEmail(email string)(bool , error){
	db := config.SetupDB()
	defer db.Close()
	
	var exist string

	err := db.QueryRow("SELECT email FROM users WHERE email = $1" , email).Scan(&exist)

	if err != nil{
		if err == sql.ErrNoRows{
			return false , nil
		}
		return false , err
	}

	return true , nil
}
