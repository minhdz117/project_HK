package api

import (
	"fmt"
	"gin_API/auth"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
	interval     = 2
	burst        = 5
	requestLimit = rate.NewLimiter(rate.Every(time.Duration(interval)*time.Second), burst)
)

type Login struct {
	USERNAME string `json:"username"`
	PASSWORD string `json:"password"`
}

func rateLimiter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !requestLimit.Allow() {
			ctx.JSON(http.StatusTooManyRequests, gin.H{
				"code":    http.StatusTooManyRequests,
				"message": http.StatusText(http.StatusTooManyRequests),
			})
			ctx.Abort()
		}
	}
}

func isAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("access_token")
		if token == "" {
			ctx.JSON(http.StatusTooManyRequests, gin.H{
				"code":    1,
				"message": "Ban chua dang nhap",
			})
			ctx.Abort()
			return
		}
		isAuthenticated, err := auth.ValidateToken(token)
		fmt.Println(isAuthenticated)
		if !isAuthenticated {
			ctx.JSON(http.StatusTooManyRequests, gin.H{
				"code":    1,
				"message": err,
			})
			ctx.Abort()
			return
		}
	}
}
func isAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("access_token")
		if token == "" {
			ctx.JSON(http.StatusTooManyRequests, gin.H{
				"code":    1,
				"message": "Ban chua dang nhap",
			})
			ctx.Abort()
			return
		}
		isAuthenticated, err := auth.ValidateToken(token)
		fmt.Println(isAuthenticated)
		if !isAuthenticated {
			ctx.JSON(http.StatusTooManyRequests, gin.H{
				"code":    1,
				"message": err,
			})
			ctx.Abort()
			return
		}
	}
}

func login(ctx *gin.Context) {
	var login Login
	ctx.BindJSON(&login)
	token, err := auth.GenerateToken(login.USERNAME, login.PASSWORD)
	if token == "" || err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "Tai khoan hoac mat khau khong chinh xac",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":         0,
		"access_token": token,
		"message":      "Dang nhap thanh cong",
	})
}

// func register(ctx *gin.Context) {

// }

func InitAPI() *gin.Engine {
	app := gin.Default()
	app.Use(rateLimiter())
	app.POST("/login", login)
	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	app.Use(isAuth())
	app.GET("/auth", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Authenticated",
		})
	})
	app.GET("/home", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "home",
		})
	})
	return app
}
