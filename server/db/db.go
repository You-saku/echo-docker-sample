package db

import (
	// .envを使うため
	"os"
	 // gorm
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

	"app/server/models" // モデルを定義してる自作パッケージ
)

// Todo: DB接続そのものを関数にしたいな
// Todo: マイグレーションは別でやりたい
func Migrate() {
	user := os.Getenv("DB_USERNAME")
    pass := os.Getenv("DB_PASSWORD")
    protocol := "tcp(db:3306)" // 文法 => tcp(コンテナ名:ポート番号) ハマるポイント
    dbName := os.Getenv("DB_NAME")

	connect := user + ":" + pass + "@" + protocol + "/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.User{}, &models.Todo{})
}

// Todo: 後で消す(存在価値が薄れているため)
func AddRecord() {
	user := os.Getenv("DB_USERNAME")
    pass := os.Getenv("DB_PASSWORD")
    protocol := "tcp(db:3306)" // 文法 => tcp(コンテナ名:ポート番号) ハマるポイント
    dbName := os.Getenv("DB_NAME")

	connect := user + ":" + pass + "@" + protocol + "/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	db.Create(&models.User{Name: "capybara", Email: "sample@example.com", Age: 20})
}
