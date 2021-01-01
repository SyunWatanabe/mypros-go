package main

import (
	"log"
	"mypros-go/utils"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", utils.RootHandler)

	e.GET("/login/github", utils.GithubLoginHandler)

	e.GET("/login/github/callback", utils.GithubCallbackHandler)

	e.GET("/mypage", utils.MypageHandler)
    
	e.Logger.Fatal(e.Start(":8000"))
}
