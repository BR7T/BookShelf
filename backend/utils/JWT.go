package utils

import "github.com/golang-jwt/jwt/v5"

type JWTCreator interface{
	ToJwtClaims() jwt.Claims
}

func CriaJWT(data JWTCreator , secretKey string)(string , error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256 , data.ToJwtClaims())
	return  token.SignedString([]byte(secretKey))
}

func VerifyJWT(jwt string , secretKey string){

}