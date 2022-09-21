package main

import (
	"demo/internal/web"
	"demo/user_service"
)

func main() {
	router := web.NewRouter()
	router.GET("/users/:id", user_service.GetAccountByIdHandler())
	web.ServeHttp(":4000", "user_service", router)
}
