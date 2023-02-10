package env 

import (
	// .envを読み込むため
	"github.com/joho/godotenv"

	// ログ
	"log"	
)

func Load() {
	err := godotenv.Load(".env") // ここに.envファイルを入れればよい(指定されたものが読み込まれる) defaultは.env
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
