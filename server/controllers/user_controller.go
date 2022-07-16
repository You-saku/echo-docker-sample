package controllers

import (
	// echo
	"net/http"
	"github.com/labstack/echo/v4"

	"app/server/models" // モデルを定義してる自作パッケージ
	"app/server/db"
)

func CreateUser(c echo.Context) error {
	db := db.Connect()
	user := &models.User{

	}
	if err := c.Bind(u); err != nil {
		return err
	}


	return c.JSON(http.StatusCreated, user)
}

func GetUsers(c echo.Context) error {

}

func GetUser(c echo.Context) error {

}

func UpdateUser(c echo.Context) error {

}

func DeleteUser(c echo.Context) error {

}
