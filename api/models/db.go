package models

import (
	// .envを使うため
	"os"
	 // gorm
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

)

// Todo:
// DBとの接続方法
// 現状Connect()を多用しているので1回1回接続してる気がする...

// DBとの接続を行いgormを使用できるようにする
func Connect() (db *gorm.DB){
	user := os.Getenv("DB_USERNAME")
    pass := os.Getenv("DB_PASSWORD")
    protocol := "tcp(db:3306)" // 文法 => tcp(コンテナ名:ポート番号) ハマるポイント
    dbName := os.Getenv("DB_NAME")

	connect := user + ":" + pass + "@" + protocol + "/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return
}

// Todo: マイグレーションはgorm以外でやりたい
func Migrate() {
	db := Connect()
	db.AutoMigrate(&User{}, &Todo{})
}

// DB挿入のサンプル関数
func AddRecord() {
	db := Connect()
	db.Create(&User{Name: "capybara", Email: "sample@example.com", Age: 20})
}
