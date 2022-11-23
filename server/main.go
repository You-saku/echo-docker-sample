package main

import (
	// 自作packageのimport
	// go.modにあるmodule名+パス(ハマったポイント)
//	"app/server/models" // DB操作一覧
	"app/server/env" // .envファイルを読み込む
	"app/server/route" // ルーティングの定義
)

func main() {
	env.Load() // 関数名は大文字じゃないと呼び出せない(ハマるポイント)
//	models.Migrate()
	route.Routing()
}
