package models

import (
	// gorm
    "gorm.io/gorm"
)

// ORM用
// 構造体名は単数で => db上では複数形になる
// カラム名は頭文字が大文字じゃないといけない(ハマるポイント)
type Todo struct {
	gorm.Model
	Id uint
	UserID uint // 自動で外部キーになる(外部キー上書きの方法はドキュメントをみてね)
	Content string
	IsFinished bool
}
