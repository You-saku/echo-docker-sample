package models

import (
	// .envを使うため
	"os"
	 // gorm
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

)

// 外部で使う場合は「:=」で変数宣言してはいけない
// https://go-tour-jp.appspot.com/basics/10
// > なお、関数の外では、キーワードではじまる宣言( var, func, など)が必要で、 := での暗黙的な宣言は利用できません。 
var db *gorm.DB
var err error

// DBとの接続を行いgormを使用できるようにする
func Init() {
	user := os.Getenv("DB_USERNAME")
    pass := os.Getenv("DB_PASSWORD")
    protocol := "tcp(db:3306)" // 文法 => tcp(コンテナ名:ポート番号) ハマるポイント
    dbName := os.Getenv("DB_NAME")

	dsn := user + ":" + pass + "@" + protocol + "/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
