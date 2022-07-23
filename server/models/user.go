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
	gorm.Model // これを入れるとcreated_atとかupdated_atを自動でなんとかしてくれる。しかし、論理削除になる
	Todos   []Todo // これでhasmany
}

// これ良いのか？
func GetAllUsers() []User {
	var users []User
	db := Connect()
	db.Find(&users)
	return users
}

// 構造体にメソッドを持たせる書き方
func (u *User) Create() (tx *gorm.DB) {
	db := Connect()
	return db.Create(&u)
}

func (u *User) Show(id uint) (tx *gorm.DB) {
	db := Connect()
	return db.Where("id = ?", id).First(&u)
}

func (u *User) Update(id uint) (tx *gorm.DB) {
	db := Connect()
	return db.Where("id = ?", id).Updates(&u)
}

func (u *User) Delete(id uint) (tx *gorm.DB) {
	db := Connect()
	return db.Where("id = ?", id).Delete(&u) // 論理削除になった
}
