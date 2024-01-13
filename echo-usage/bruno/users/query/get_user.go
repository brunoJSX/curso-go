package query

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetUserHandler(c echo.Context) error {
	db, _ := initDB()
	defer db.Close()

	var u User
	row := db.QueryRow("select id, name from users where id = ?", c.Param("id"))
	err := row.Scan(&u.ID, &u.Name)
	if err != nil {
		return c.JSON(http.StatusNotFound, ResourceNotFound{Status_Code: 404, Message: "User not found!"})
	}

	return c.JSON(http.StatusOK, u)
}
