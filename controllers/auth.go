// Package controllers - contains all the controllers
// @SubApi User [/auth]
// @SubApi Allows you access to different features of the users , login , registration, etc [/auth]
package controllers

import (
	"fmt"
	"net/http"
	"os"
	"site-checker/db"
	"site-checker/models"
	"site-checker/utils"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// Login - get auth token with username and password
// @Title Login User
// @Description Login User and retrieve a token for the authentication
// @Accept json
// @Param userId path int true &quot;User ID&quot;
// @Success 200 {object} string &quot;Success&quot;
// @Failure 401 {object} string &quot;Access denied&quot;
// @Failure 404 {object} string &quot;Not Found&quot;
// @Resource /auth
// @Router /v1/auth/:userId.json [get]
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

// Registration - create new user
func Registration(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	// encrypt password
	password, err := utils.EncryptPassword(os.Getenv("SALT"), c.FormValue("password"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	u.Password = password

	fmt.Println(password)

	if res := db.DBCon.Create(u); res.Error != nil {
		return c.JSON(http.StatusBadRequest, res.Error)
	}
	return c.JSON(http.StatusCreated, u)
}
