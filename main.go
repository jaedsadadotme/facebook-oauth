package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "hello"})
	})
	e.GET("/oauth/facebook", func(c echo.Context) error {
		client_id := os.Getenv("CILENT_ID")
		uri := fmt.Sprintf("https://www.facebook.com/v8.0/dialog/oauth?client_id=%&redirect_uri=http://localhost:1323/&scope=email&state=null", client_id)
		return c.Redirect(http.StatusPermanentRedirect, uri)
	})
	e.GET("/oauth/facebook/userInfo", func(c echo.Context) error {
		// get Bearer Token
		token, err := getTokenBearer(c.QueryParam("code"))
		if err.ErrorCode == 400 {
			return c.JSON(400, err)
		}
		result := getUserInfo(token.AccessToken)
		return c.JSON(200, result)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
