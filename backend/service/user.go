package service

import (
	"fmt"
	"time"

	"github.com/BR7T/BookShelf/config"
)

type User struct {
    ID        int64
    Name      string
    Email     string
    CreatedAt time.Time
}

func GetAllUsers()([]User , error){
	db := config.SetupDB()
	// search all users
	rows , err := db.Query("SELECT iduser , username , email , createdAt FROM users")
	if err != nil{
		return nil , fmt.Errorf("failed to search users %w" , err)
	}

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID , &user.Name , &user.Email , &user.CreatedAt)
		if err != nil{
			return nil , fmt.Errorf("failed to search users %w" , err)
		}
		users = append(users, user)
	} 
	rows.Close()

	return  users , nil
}