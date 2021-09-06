package main

import (
	"fmt"
	"tudai2021/services"
)

func main() {
	var userSrv services.UserService             //interfaz
	userSrv, err := services.NewUserServiceAWS() //devuelve una implementacion

	if err != nil {
		panic(err)
	}

	u, err := userSrv.Login("jppAWS", "654321")
	if err != nil {
		panic(err)
	}
	fmt.Println(u)
}
