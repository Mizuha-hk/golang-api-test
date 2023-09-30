package handler

import (
	"net/http"
	"strconv"
	"time"

	"golang-api/api/models" // ユーザーモデルのパッケージを適切にインポートしてください

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	// ユーザーの検証
	user, err := models.GetUserByNameAndPassword(req.Name, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
	}

	// JWTトークンの生成
	token, err := generateJWTToken(strconv.Itoa(user.ID)) // IDをstringに変換
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Could not generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func generateJWTToken(userID string) (string, error) {
	// シークレットキーを設定
	secretKey := "your_secret_key" // 実際のアプリケーションでは、安全な方法でキーを管理してください
	// トークンの有効期限を設定
	expirationTime := time.Now().Add(24 * time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     expirationTime.Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
