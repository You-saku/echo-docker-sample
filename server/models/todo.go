package models

import (
	// gorm
    "gorm.io/gorm"
)

// ORM用
// 構造体名は単数で => db上では複数形になる
// カラム名は頭文字が大文字じゃないといけない(ハマるポイント)
type Todo struct {
	ID int
	UserID int // 外部キーになる(db上ではuser_idになる)
	Content string
	IsFinished bool
	gorm.Model
}
