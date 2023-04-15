package models

import (
	// gorm
	"gorm.io/gorm"
)

// ORM用
// 構造体名は単数で => db上では複数形になる
// カラム名は頭文字が大文字じゃないといけない(ハマるポイント)
type User struct {
	ID int `param:"id"`
	Name  string `form:"name" json:"name" validate:"required"`
	Email string `form:"email" json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Age int `form:"age" json:"age" validate:"gte=0,lte=100"`
	gorm.Model // これを入れるとcreated_atとかupdated_atを自動でなんとかしてくれる。しかし、論理削除になる
	Todos   []Todo `gorm:"foreignkey:UserID"` // これでhasmany
}

func GetAllUsers() []User {
	var users []User
	db.Find(&users)
	return users
}

// 構造体にメソッドを持たせる書き方
func (u *User) Create() (tx *gorm.DB) {
	return db.Create(&u)
}

func (u *User) Show(id uint) (tx *gorm.DB) {
	// リレーション先を取得する
	return db.Where("id = ?", id).First(&u).Preload("Todos").Find(u)
}

func (u *User) Update(id uint) (tx *gorm.DB) {
	return db.Where("id = ?", id).Updates(&u)
}

func (u *User) Delete(id uint) (tx *gorm.DB) {
	return db.Where("id = ?", id).Delete(&u) // 論理削除になった
}

func (u *User) Verify(email string, password string) (tx *gorm.DB) {
	return db.Where("email = ? AND password = ?", email, password).First(&u)
}
