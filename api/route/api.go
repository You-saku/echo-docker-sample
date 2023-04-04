package route

import (
	// echo
	"net/http"
	"github.com/labstack/echo/v4"

	// modelを読み込む
	"app/api/models"
	// controllerを読み込む
	"app/api/controllers"
	// バリデーション
	"app/api/validates"

	// ミドルウェア導入
	"github.com/labstack/echo/v4/middleware"
	"app/api/token"
)

func Routing() {
	e := echo.New() // echoのインスタンスを作成
	models.Init() // DBとの接続を行う

	// CORS設定
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// Reactのフロント側, swagger-uiの2つを許可
		AllowOrigins: []string{"http://localhost:3001", "http://localhost:8000"},
	}))

	// サンプルのルーティング
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Echo!")
	})
	e.GET("/sample", sample)

	// バリデーションを設定
	e.Validator = validates.SetValidator()

	// prefixつきrouting
	api := e.Group("/api")
	api.GET("/ok", func(c echo.Context) error {
		return c.String(http.StatusOK, "api")
	})

	// ログイン(jwtを返す)
	api.POST("/login", controllers.Login)

	// ログインしてたら叩けるエンドポイントを作る
	login := api

	// 認証設定
	config := middleware.JWTConfig{
		Claims:     &token.JwtCustomClaims{},
		SigningKey: []byte(token.Key),
	}
	// メモ
	// このJWTWithConfigを使うと間違ったJWTなら401, Authorization headerがなければ400になる
	login.Use(middleware.JWTWithConfig(config))

	// userのCRUD
	login.POST("/users", controllers.CreateUser)
	login.GET("/users", controllers.GetUsers)
	login.GET("/users/:id", controllers.GetUser)
	login.PUT("/users/:id", controllers.UpdateUser)
	login.DELETE("/users/:id", controllers.DeleteUser)

	e.Logger.Fatal(e.Start(":80"))
}

// ルーティングに使う関数
func sample(c echo.Context) error {
	return c.String(http.StatusOK, "Sample.")
}
