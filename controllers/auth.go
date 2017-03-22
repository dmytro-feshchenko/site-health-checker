package controllers

import (
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Login - get auth token with username and password
func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "myuser" && password == "1234" {
		// create jwt token
		token := jwt.New(jwt.SigningMethodHS256)
		tokenExpirationTime := time.Now().Add(time.Hour * 72)

		// set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Dmitryi Feschenko"
		claims["admin"] = true
		claims["exp"] = tokenExpirationTime.Unix()

		// generate encoded token and send it to response
		t, err := token.SignedString([]byte(os.Getenv("SECRET")))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token_type":   "Bearer",
			"expires_in":   tokenExpirationTime.String(),
			"access_token": t,
		})
	}

	return echo.ErrUnauthorized
}
