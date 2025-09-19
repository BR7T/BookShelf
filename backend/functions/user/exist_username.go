package functions

import (
	"database/sql"

	"github.com/BR7T/BookShelf/config"
)

func ExistUsername(username string)(bool , error){
	db := config.SetupDB()
	defer db.Close()

	var exist string

	err := db.QueryRow("SELECT username FROM users WHERE username = $1" , username).Scan(&exist)

	if err != nil{
		if err == sql.ErrNoRows{
			return false , nil
		}
		return false , err
	}

	return true , nil
}

