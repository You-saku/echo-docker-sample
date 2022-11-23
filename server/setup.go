/** migrationやレコード挿入を行う */
package main

import (
	"app/server/env" // .envファイルを読み込む
	"app/server/models" // DB操作一覧
)

/** マイグレーションする */
func main() {
	env.Load()
	models.Migrate()
	models.AddRecord()
}
