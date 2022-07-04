package main

import (
	"errors"
	"fmt"
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

func ValidateToken() (bool, string) {

	signingKey := []byte(secret)
	// userClaims := jwtclaim{
	// 	Username: "minhdz117",
	// 	Role:     "ADMIN",
	// 	Exp:      time.Now().Add(time.Hour * 24).Unix(),
	// }
	// token, err := jwt.Sign(jwt.HS256, signingKey, userClaims)
	// if err != nil {
	// 	return false, "Loi tao token"
	// }
	// fmt.Println(string(token))
	t := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im1pbmhkejExNyIsInJvbGUiOiJBRE1JTiIsImV4cCI6MTY1NjQ3MTUxN30.3Rp2KpnChLe8jp2MyIZ3Y8vu2oqcTQpyBcFVvzHP1Zk"
	token := []byte(t)
	verifiedToken, err := jwt.Verify(jwt.HS256, signingKey, token)
	if err != nil {
		print(err)
		return false, "Khong the giai ma"
	}
	var claims = struct {
		Token string `json:"token"`
		Exp   int64  `json:"exp"`
	}{}
	verifiedToken.Claims(&claims)
	if claims.Exp < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		println(err)
		return false, "JWT is expired"
	}
	return true, ""

}
func main() {
	status, err := ValidateToken()
	fmt.Println(status)
	fmt.Println(err)
}
