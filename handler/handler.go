package handler

import (
	"fmt"
	"golang101/model"
	"time"

	"github.com/dgrijalva/jwt-go"

	"net/http"

	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

type simpleResponse struct {
	Username string
	Point    float64
}

type loginRequest struct {
	Username string
	Password string
}

type user struct {
	Username string
}
type JWT struct {
	Name string
	Role bool
}
type welcome struct {
	Massage string
}

var bearer string

//ServeHomeReturnJSON is render homepage
func ServeHomeReturnJSON(c echo.Context) error {
	jsonObj := simpleResponse{
		"nine", 10.2,
	}
	return c.JSON(http.StatusOK, jsonObj)
}

//ServeHomePath is to render homepage
func ServeHomePath(c echo.Context) error {
	res := welcome{
		"Welcome",
	}
	return c.JSON(http.StatusOK, res)
}

//ServeHomePathReceivePostBody is render homepage
func ServeHomePathReceivePostBody(uls model.UserLoginLogDB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody loginRequest
		fmt.Println("Before: ", reqBody)
		err := c.Bind(&reqBody)
		if err != nil {
			fmt.Println("something worng")
		}
		fmt.Println("After: ", reqBody)
		userLoginLog := &model.UserLoginLog{Username: reqBody.Username, Password: reqBody.Password}
		err = uls.Create(userLoginLog)
		if err != nil {
			fmt.Println("error when crateing", err)
		}
		return c.JSON(http.StatusOK, reqBody)
	}
}

//UserSignin is render homepage
func UserSignin(uls model.UserLoginLogDB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody loginRequest
		err := c.Bind(&reqBody)
		if err != nil {
			fmt.Println("something wrong")
		}
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   1200 * 1,
			HttpOnly: true,
		}
		sess.Values["user"] = reqBody.Username
		sess.Values["pass"] = reqBody.Password
		sess.Save(c.Request(), c.Response())
		return c.JSON(http.StatusOK, reqBody)
	}
}

//GetSession is show user login
func GetSession(c echo.Context) error {
	sess, _ := session.Get("session", c)
	name := sess.Values["user"]
	pass := sess.Values["pass"]
	obj := make(map[string]interface{})
	obj["name"] = name
	obj["pass"] = pass
	return c.JSON(http.StatusOK, obj)
}

//QueryDataUser is show user
func QueryDataUser(c echo.Context) error {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	var users []model.Product
	db.Find(&users)

	return c.JSON(http.StatusOK, users)
}

//Accessible is Accessible
func Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

//Restricted is Accessible
func Restricted(c echo.Context) error {

	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*model.JwtCustomClaims)
	name := claims.Name
	role := claims.Admin
	return c.JSON(http.StatusOK, JWT{
		name,
		role,
	})
}

func Login(c echo.Context) error {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	var reqBody model.UserLoginLog
	err = c.Bind(&reqBody)
	if err != nil {
		return echo.ErrUnauthorized
	}
	// Throws unauthorized error
	result := db.Select("Username").Where("username = ?", reqBody.Username).Where("password = ?", reqBody.Password).First(&reqBody)
	if result.Error != nil {
		return echo.ErrUnauthorized
	}
	// Set custom claims
	claims := &model.JwtCustomClaims{
		reqBody.Username,
		true,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}
	bearer = "Bearer " + t
	c.Request().Header.Set("Authorization", bearer)
	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
