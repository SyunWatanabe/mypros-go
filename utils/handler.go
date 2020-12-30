package utils

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/labstack/echo/v4"
)

// RootHandler root
func RootHandler(c echo.Context) error {
	const tag = `<a href="/login/github/">LOGIN</a>`

	c.Response().Header().Set("Content-Type", "application/json")
	_, err := c.Response().Write([]byte(tag))
	if err != nil {
		log.Panic("rootHandler failed", err)
	}
	return err
}

// MypageHandler /mypage
func MypageHandler(c echo.Context) error {
	return c.String(http.StatusOK, "mypage")
}

var token = OauthStateString()

// GithubLoginHandler /login/github
func GithubLoginHandler(c echo.Context) error {
	url := Config().AuthCodeURL(token)

	return c.Redirect(http.StatusTemporaryRedirect, url)
}

// GithubCallbackHandler /login/github/callback
func GithubCallbackHandler(c echo.Context) error {
	state := c.FormValue("state")
	if state != token {
		fmt.Printf("invalid oauth state, expected %s', got %s\n", token, state)
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	code := c.FormValue("code")
	token, err := Config().Exchange(context.Background(), code)
	if err != nil {
		fmt.Printf("Config().Exchange() failed with '%s, error: %s'\n", token, err)
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	oauthClient := Config().Client(context.Background(), token)
	client := github.NewClient(oauthClient)

	user, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		fmt.Printf("client.Users.Get() failed with '%s'\n", err)
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}

	fmt.Printf("Logged in as Github user: %s\n", *user.Name)
  return c.Redirect(http.StatusMovedPermanently, "/mypage")
}
