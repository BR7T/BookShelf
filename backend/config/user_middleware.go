package config

import "net/http"

func MiddlewareUser(next http.Handler)http.Handler{
	return next
}

func MiddlewareAdmin(next http.Handler)http.Handler{
	return  next
}