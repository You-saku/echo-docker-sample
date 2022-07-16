package models

import (
	// gorm
	"gorm.io/gorm"
)

// ORM用
// 構造体名は単数で => db上では複数形になる
// カラム名は頭文字が大文字じゃないといけない(ハマるポイント)
type User struct {
	ID uint
	Name  string `form:"name" json:"name"`
	Email string `form:"email" json:"email"`
	Age uint `form:"age" json:"age"`
	gorm.Model
	Todos   []Todo // これでhasmany
}

func (u *User) Create() (tx *gorm.DB) {
	db := Connect()
	return db.Create(&u)
}
