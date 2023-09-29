package main

import (
	"golang-api/api/router"
	"golang-api/config"
	"golang-api/db"

	"github.com/labstack/echo/v4"
)

func main() {
    // 設定を読み込み
    config := config.GetConfig()

    // データベース接続
    db.Connect(config.Database)
    
    // Echoインスタンスの作成
    e := echo.New()
    
    // ルーター設定
    router.SetupRouter(e)
    
    // サーバー起動
    e.Start(":" + config.Server.Port)
}
