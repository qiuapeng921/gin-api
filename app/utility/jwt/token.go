package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var (
	jwtSecret = os.Getenv("JWT_SECRET")
)

type MapClaims struct {
	Id       uint   `json:"id"`
	Account  string `json:"account"`
	Category string `json:"category"`
	jwt.StandardClaims
}

// GenerateToken generate tokens used for auth
func GenerateToken(id uint, account, category string) (string, int64, error) {
	nowTime := time.Now()
	// token 有效期 24小时
	expireTime := nowTime.Add(24 * time.Hour)

	claims := MapClaims{
		id,
		account,
		category,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(jwtSecret))

	return token, claims.ExpiresAt, err
}

// ParseToken parsing token
func ParseToken(token string) (*MapClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MapClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
