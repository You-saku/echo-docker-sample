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
	user := new(models.User)
	c.Bind(user) // これ最強
	user.Create()

	return c.JSON(http.StatusCreated, user)
}

func GetUsers(c echo.Context) error {
	users := models.GetAllUsers()

	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	user := models.User{}
	user.Show(id)

	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)

	user := new(models.User)
	c.Bind(user)
	user.Update(id)

	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	user := models.User{}
	user.Delete(id)

	return c.NoContent(http.StatusNoContent)
}
