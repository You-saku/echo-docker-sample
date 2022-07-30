package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"

	"app/server/token"
)

// Todo DBに存在するユーザーだけログインさせる
func Login(c echo.Context) (err error) {
	jwtClaim := new(token.JwtCustomClaims)

	c.Bind(jwtClaim)
	if err := c.Validate(jwtClaim); err != nil {
		return err
	}
	t, err := token.CreateJwt(jwtClaim.Email, jwtClaim.Password)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

