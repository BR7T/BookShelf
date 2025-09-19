package models

import "time"

type User struct {
    ID        int64
    Name      string
    Email     string
    CreatedAt time.Time
}

type UserRegister struct{
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
	UserName string `json:"username"`
}

type UserLogin struct{
	Email string `json:"email"`
	Password string `json:"password"`
}

type UserPassword struct{
	Email string
	Password string
	Role string
}