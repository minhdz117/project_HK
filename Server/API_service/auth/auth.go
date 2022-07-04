package auth

import (
	"gin_API/database"
	"time"

	"github.com/kataras/jwt"
)

var (
	secret = "minhdz117"
)

type jwtclaim struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Exp      int64  `json:"exp"`
}

func GenerateToken(username string, password string) (string, error) {
	user, err := database.GetUserByUsername(username)
	if err != nil {
		return "", err
	}
	if user.Password != password {
		return "", err
	}

	signingKey := []byte(secret)
	userClaims := jwtclaim{
		Username: user.Username,
		Role:     user.Role,
		Exp:      time.Now().Add(time.Hour * 24).Unix(),
	}

	token, err := jwt.Sign(jwt.HS256, signingKey, userClaims)
	tokenString := string(token[:])
	return tokenString, err

}

func ValidateToken(token string) (bool, string) {
	signingKey := []byte(secret)
	tokenString := []byte(token)
	verifiedToken, err := jwt.Verify(jwt.HS256, signingKey, tokenString)
	if err != nil {
		return false, "Khong the giai ma"
	}
	var claims jwtclaim
	verifiedToken.Claims(&claims)
	if claims.Exp < time.Now().Local().Unix() {
		return false, "JWT is expired"
	}
	return true, ""

}
