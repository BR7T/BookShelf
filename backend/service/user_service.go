package service

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/BR7T/BookShelf/config"
	"github.com/BR7T/BookShelf/utils"
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


type UserRegister struct{
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
	UserName string `json:"username"`
}

func RegisterUser(user UserRegister) (bool, error) {
    existEmail, err := existEmail(user.Email)
    if err != nil {
        return false, fmt.Errorf("error checking email: %w", err)
    }
    if existEmail {
        return false, fmt.Errorf("email already exists")
    }

    existUsername, err := existUsername(user.UserName)
    if err != nil {
        return false, fmt.Errorf("error checking username: %w", err)
    }
    if existUsername {
        return false, fmt.Errorf("username already exists")
    }

    hash, err := utils.GenerateHashArgon2(user.Password)
    if err != nil {
        return false, fmt.Errorf("failed to hash password: %w", err)
    }

    db := config.SetupDB()
    defer db.Close()

    _, err = db.Exec(
        "INSERT INTO users (password, username, name, email) VALUES ($1, $2, $3, $4)",
        hash, user.UserName, user.Name, user.Email,
    )
    if err != nil {
        return false, fmt.Errorf("failed to insert user: %w", err)
    }

    return true, nil
}

type UserLogin struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(login UserLogin)(bool , error){
	DatabaseUser , err := getHashPassword(login.Email)
	if err != nil{
		return false , err
	}
	valid , err := utils.ValidArgon2(login.Password , DatabaseUser.Password)
	if err != nil{
		return false , err
	}
	if valid{
		return true , nil	
	}else{
		return false , nil
	}
}


type UserPassword struct{
	Email string
	Password string
}
func getHashPassword(email string)(*UserPassword , error){
	var user UserPassword

	db := config.SetupDB()
	err := db.QueryRow("SELECT email , password FROM users WHERE email = $1" , email).Scan(&user.Email , &user.Password)
	if err != nil{
		return nil , fmt.Errorf("erro ao fazer a consulta sql")
	}

	return &user , nil
}


func existUsername(username string)(bool , error){
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

func existEmail(email string)(bool , error){
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
