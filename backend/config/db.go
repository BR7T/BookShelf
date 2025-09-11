package config

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

func SetupDB() *sql.DB{
	godotenv.Load()
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	host := os.Getenv("DB_HOST")
	connectionS := fmt.Sprintf("host=%s port =%s user=%s password=%s dbname=%s sslmode=disable",host , port , user , password , database)

	db , err := sql.Open("postgres" , connectionS)

	if err != nil{
		fmt.Print(err)
		return nil
	}
	
	return db
}