package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var (
	// 自定义 对称加密密钥（HS256 用）
	jwtKey = []byte("victree")

	// 自定义 token 过期时间（3 天）
	tokenExpire = time.Hour * 3 * 24
)

// 自定义 JWT Claims 结构体
type VictreeClaims struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

/* 生成 JWT Token */
func GenerateToken(userId int, username string) (string, error) {
	// 1、 创建 Claims
	claims := &VictreeClaims{
		UserId:   userId,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpire)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                  // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                  // 生效时间
			Issuer:    "victree",                                       // 签发人
		},
	}

	// 2、创建 Token (指定算法和 Claims )
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 3、签名（使用密钥）
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

/* 验证 JWT Token */
func ValidateToken(tokenString string) (*VictreeClaims, error) {
	// 1、解析 token
	token, err := jwt.ParseWithClaims(tokenString, &VictreeClaims{}, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})

	if err != nil {
		// 常见错误：令牌过期（ErrTokenExpired）、签名错误（ErrSignatureInvalid）等
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("令牌已过期")
		} else {
			return nil, errors.New("无效的令牌")
		}
	}

	// 2、验证 token
	if claims, ok := token.Claims.(*VictreeClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, errors.New("无效的 token")
	}
}
