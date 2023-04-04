package main

import (
	// 自作packageのimport
	// go.modにあるmodule名+パス(ハマったポイント)

	"app/api/env" // .envファイルを読み込む
	"app/api/route" // ルーティングの定義
)

func main() {
	env.Load() // 関数名は大文字じゃないと呼び出せない(ハマるポイント)
	route.Routing()
}
