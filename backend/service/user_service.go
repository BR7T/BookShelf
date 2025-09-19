package service

import (
	"fmt"
	"os"

	"github.com/BR7T/BookShelf/config"
	functions "github.com/BR7T/BookShelf/functions/user"
	"github.com/BR7T/BookShelf/models"
	"github.com/BR7T/BookShelf/utils"
)

func RegisterUser(user models.UserRegister) (bool, error ) {
    existEmail, err := functions.ExistEmail(user.Email)
    if err != nil {
        return false, fmt.Errorf("error checking email: %w", err)
    }
    if existEmail {
        return false, fmt.Errorf("email already exists")
    }

    existUsername, err := functions.ExistUsername(user.UserName)
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

func LoginUser(login models.UserLogin)(bool , error){
	DatabaseUser , err := functions.GetHashPassword(login.Email)
	if err != nil{
		return false , err
	}
	valid , err := utils.ValidArgon2(login.Password , DatabaseUser.Password)
	if err != nil{
		return false , err
	}
	if valid{
        // se for válido gera JWT
        userJWT := &models.UserJWT{
            Email: DatabaseUser.Email,
            Role: DatabaseUser.Role,
        }
        jwtToken , err := utils.CriaJWT(userJWT , os.Getenv("JWT_PASSWORD"))

        if err!=nil{
            return false , err
        }
        
        fmt.Print(jwtToken)

		return true , nil
	}else{
		return false , nil
	}
}






