package command

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CreateUser(c echo.Context, db *sql.DB) error {
	defer db.Close()

	var u User
	if err := c.Bind(&u); err != nil {
		return err
	}

	stmt, _ := db.Prepare("insert into users(name) values(?)")
	stmt.Exec(u.Name)

	return c.JSON(http.StatusCreated, SuccessMessage{Status_Code: 201, Message: "User created successfully!"})
}
