package main

import (
	"github.com/labstack/echo/v4"
)

type ResourceNotFound struct {
	Status_Code int    `json:"status_code"`
	Message     string `json:"message"`
}

type SuccessMessage struct {
	Status_Code int    `json:"status_code"`
	Message     string `json:"message"`
}

func main() {
	e := echo.New()

	e.GET("/users/:id", query.getUserHandler)
	e.POST("/users/", command.CreateUser)
	e.DELETE("/users/:id", command.DeleteUser)

	e.Logger.Fatal(e.Start(":3001"))
}
