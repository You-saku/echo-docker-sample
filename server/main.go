package main

import (
	// echo
	"net/http"
	"github.com/labstack/echo/v4"

	// .envを使うため
	"github.com/joho/godotenv"
	"os"

	 // gorm
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

	// ログ
	"log"
	
)

// ORM用
// 構造体名は単数で => db上では複数形になる
// カラム名は頭文字が大文字じゃないといけない(ハマるポイント)
type User struct {
	gorm.Model
	Name  string
	Email string
	Age uint
}

func main() {
	loadEnv()
	// migrateDB()
 	// addRecord()

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


	e.Logger.Fatal(e.Start(":80"))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func migrateDB() {
	user := os.Getenv("DB_USERNAME")
    pass := os.Getenv("DB_PASSWORD")
    protocol := "tcp(db:3306)" // 文法 => tcp(コンテナ名:ポート番号) ハマるポイント
    dbName := os.Getenv("DB_NAME")

	connect := user + ":" + pass + "@" + protocol + "/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
}

func addRecord() {
	user := os.Getenv("DB_USERNAME")
    pass := os.Getenv("DB_PASSWORD")
    protocol := "tcp(db:3306)" // 文法 => tcp(コンテナ名:ポート番号) ハマるポイント
    dbName := os.Getenv("DB_NAME")

	connect := user + ":" + pass + "@" + protocol + "/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&User{Name: "capybara", Email: "sample@example.com", Age: 20})
}

// ルーティングに使う関数
func sample(c echo.Context) error {
	return c.String(http.StatusOK, "Sample.")
}
