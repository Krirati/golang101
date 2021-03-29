package middleware

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

type ResponseFail struct {
	Code    string
	Message string
}

type LoggerConfig struct {
	Format string `json:"format"`
	Output io.Writer
}

//GateKeeper si should logs all call to echo path
func GateKeeper(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("somthing get call")
		return next(c)
	}
}

//SessionAuth is check login
func SessionAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("SessionAuth get call")
		sess, _ := session.Get("session", c)
		name := sess.Values["user"]
		pass := sess.Values["pass"]
		if name != nil {
			obj := make(map[string]interface{})
			obj["name"] = name
			obj["pass"] = pass
			return next(c)
		} else {
			fmt.Println("Please login!")
			var res ResponseFail
			res.Code = "Please login!"
			res.Message = "Please login!"
			return c.JSON(http.StatusOK, res)
		}
	}
}

//LoggerWithConfig is log
func LoggerWithConfig(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		log.Println(time.Since(start), c.Request().Method, c.Request().RequestURI)
		return next(c)
	}
}

//CheckJWT is check role
func CheckJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request().Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoia3JpcmF0aSIsImFkbWluIjp0cnVlLCJleHAiOjE1ODgwMDMwNDV9._8EUtbvlR_qGsl0VSUWROttGeKzl4lG8s_I2e43J0dw")
		if c.Request().Header.Get("Authorization") != "" {
			return next(c)
		} else {
			return c.JSON(http.StatusOK, echo.ErrUnauthorized)
		}
	}
}
