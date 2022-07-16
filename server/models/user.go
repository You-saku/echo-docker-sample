package models

import (
	// gorm
	"gorm.io/gorm"
)

// ORM用
// 構造体名は単数で => db上では複数形になる
// カラム名は頭文字が大文字じゃないといけない(ハマるポイント)
type User struct {
	gorm.Model
	Id uint
	Name  string
	Email string
	Age uint
	Todos   []Todo // これでhasmany
}
