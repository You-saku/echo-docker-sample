package route

import (
	// echo
	"net/http"
	"github.com/labstack/echo/v4"

	// controllerを読み込む
	"app/server/controllers"
	// バリデーション
	"app/server/validates"

	// ミドルウェア
	"github.com/labstack/echo/v4/middleware"
	"app/server/token"
)

func Routing() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})

	e.POST("/", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		return c.String(http.StatusOK, "name:" + name + ", email:" + email)
	})

	e.GET("/sample", sample)

	// バリデーションを設定
	e.Validator = validates.SetValidator()

	// ログイン(jwtを返す)
	e.POST("/login", controllers.Login)

	// prefixつきrouting
	api := e.Group("/api")

	// 認証を求めるエンドポイントに設定
	config := middleware.JWTConfig{
		Claims:     &token.JwtCustomClaims{},
		SigningKey: []byte(token.Key),
	}
	api.Use(middleware.JWTWithConfig(config))

	api.GET("/ok", func(c echo.Context) error {
		return c.String(http.StatusOK, "api")
	})

	// userのCRUD
	api.POST("/user", controllers.CreateUser)
	api.GET("/users", controllers.GetUsers)
	api.GET("/users/:id", controllers.GetUser)
	api.PUT("/users/:id", controllers.UpdateUser)
	api.DELETE("/users/:id", controllers.DeleteUser)

	e.Logger.Fatal(e.Start(":80"))
}

// ルーティングに使う関数
func sample(c echo.Context) error {
	return c.String(http.StatusOK, "Sample.")
}
