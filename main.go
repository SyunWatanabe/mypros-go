package main

import(
	"net/http"
    "github.com/labstack/echo/v4"
    _ "github.com/go-sql-driver/mysql"
    _ "gorm.io/gorm"
    _ "gorm.io/driver/mysql"
    _ "github.com/valyala/fasthttp"
    _ "go.uber.org/zap"
    _ "github.com/joho/godotenv"
)

type User struct {
    Name string `json:"name"`
}

func main() {
    e := echo.New()
    u := new(User)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, u)
	})
	e.Logger.Fatal(e.Start(":8000"))
}
