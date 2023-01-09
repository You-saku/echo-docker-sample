package main

import (
	// 自作packageのimport
	// go.modにあるmodule名+パス(ハマったポイント)
	"app/server/env" // .envファイルを読み込む
	"app/server/models" // モデルを読み込む
	"app/server/route" // ルーティングの定義
)

func main() {
	env.Load() // 関数名は大文字じゃないと呼び出せない(ハマるポイント)
	models.Migrate() // マイグレーションした後に呼び出しても問題はない(データは消されない)
	route.Routing()
}
