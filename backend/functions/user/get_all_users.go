package functions

import (
	"fmt"

	"github.com/BR7T/BookShelf/config"
	"github.com/BR7T/BookShelf/models"
)

func GetAllUsers()([]models.User , error){
	db := config.SetupDB()
	// search all users
	rows , err := db.Query("SELECT iduser , username , email , createdAt FROM users")
	if err != nil{
		return nil , fmt.Errorf("failed to search users %w" , err)
	}

	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID , &user.Name , &user.Email , &user.CreatedAt)
		if err != nil{
			return nil , fmt.Errorf("failed to search users %w" , err)
		}
		users = append(users, user)
	} 
	rows.Close()

	return  users , nil
}
