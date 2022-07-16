package controllers

// Todo: Cotent-Type application-jsonでやりたい
import (
	// echo
	"net/http"
	"github.com/labstack/echo/v4"

	// 文字列を数字に変換する
	"strconv"

	"app/server/models" // モデルを定義してる自作パッケージ
)

func CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	a, _ := strconv.Atoi(c.FormValue("age"))
	age := uint(a)

	// どうやらフォームバリューからとるのが問題らしいです
	user := &models.User{
		Name: name,
		Email: email,
		Age: age,
	}
	user.Create()

	return c.JSON(http.StatusCreated, user)
}

// func GetUsers(c echo.Context) error {

// }

// func GetUser(c echo.Context) error {

// }

// func UpdateUser(c echo.Context) error {

// }

// func DeleteUser(c echo.Context) error {

// }
