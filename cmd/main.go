package main

import (
	"golang-api/api/router"
	"golang-api/config"
	"golang-api/db"
	"golang-api/api/models"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	// 設定を読み込み
	config := config.GetConfig()

	// データベース接続
	var err error
	db.DB, err = db.Connect(config.Database)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// 各モデルに対してAutoMigrateを実行
	if err := db.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
	}
	// もし他のモデルがあれば、同様に
	// if err := db.DB.AutoMigrate(&models.OtherModel{}); err != nil {
	// 	log.Fatalf("Failed to migrate the database: %v", err)
	// }

	// Echoインスタンスの作成
	e := echo.New()

	// ルーター設定
	router.SetupRouter(e)

	// サーバー起動
	e.Start(":" + config.Server.Port)
}
