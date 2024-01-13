package command

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func DeleteUser(c echo.Context) error {
	db, _ := initDB()
	defer db.Close()

	row, err := db.Exec("delete from users where id = ?", c.Param("id"))

	rowsAffected, _ := row.RowsAffected()
	if rowsAffected == 0 {
		return c.JSON(http.StatusNotFound, ResourceNotFound{Status_Code: 404, Message: "User doesn't exists"})
	}

	if err != nil {
		return c.JSON(http.StatusBadGateway, SuccessMessage{Status_Code: 500, Message: "Error on user delete"})
	}

	return c.JSON(http.StatusOK, SuccessMessage{Status_Code: 200, Message: "User deleted successfully"})
}
