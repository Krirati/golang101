package server

import (
	"fmt"
	"golang101/handler"
	"golang101/middleware"
	"golang101/model"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/line/line-bot-sdk-go/linebot"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	echoM "github.com/labstack/echo/middleware"
)

// Server is a struct for absraction
type Server struct {
	Echo           *echo.Echo
	DB             *gorm.DB
	UserLoginLogDB model.UserLoginLogDB
}

//ServerRoute is a metohd to server path
func (s *Server) ServerRoute() {
	e := s.Echo
	e.Use(middleware.LoggerWithConfig)
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.Use(echoM.CORS())

	e.GET("/", handler.ServeHomePath)
	e.GET("/json", handler.ServeHomeReturnJSON, middleware.SessionAuth)
	e.GET("/session", handler.GetSession)
	e.POST("/create-user", handler.ServeHomePathReceivePostBody(s.UserLoginLogDB), middleware.GateKeeper)
	e.POST("/signin", handler.UserSignin(s.UserLoginLogDB), middleware.GateKeeper)
	e.GET("/get-image", handler.GetImage)
	e.GET("/get-json", handler.GetResponseJson)
	// e.GET("/users", handler.QueryDataUser)
	e.GET("/acc", handler.Accessible)

	// ------- web socket------////
	e.Use(echoM.Recover())
	e.GET("/ws", handler.WsEndpoint)
	// ------- web socket------////

	e.POST("/login", handler.Login)
	config := echoM.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte("secret"),
	}
	r := e.Group("/restricted")
	r.Use(echoM.JWTWithConfig(config))
	r.GET("", handler.Restricted)
	r.GET("/user", handler.QueryDataUser)

	//// Line bot/////
	e.POST("/callback", func(c echo.Context) error {

		bot, err := linebot.New("5b3b0f7d68766d382f6e54125795fea0", "PoE1ahGSYOfMBGoSQXWqtGHeOynU+dyAQPGvXAi0/q+Hbcl3U+PwK2/6U4BDGy5LzP0TKgV6KPOAoJERQ73vp2aMkb+8kYm77OgsgirXhL5WTX7dqVQQuhGwJSIsXV84EQqkGXM2y55n4pAo0auYBwdB04t89/1O/w1cDnyilFU=")
		//////////////////
		events, err := bot.ParseRequest(c.Request())
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				c.NoContent(http.StatusBadRequest)
			} else {
				c.NoContent(http.StatusInternalServerError)
			}
			return err
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					fmt.Println(message.Text)
					if message.Text != "Test" {
						bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("BooooYahhh!")).Do()
						continue
					}
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				case *linebot.StickerMessage:
					fmt.Println(message.StickerID)
					bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("11537", "52002735")).Do()
					continue
				}
			}
		}
		return c.NoContent(http.StatusOK)
	})
}

// ServerHTTP should serve an endpoint
func (s *Server) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	s.Echo.ServeHTTP(w, r)
}

//Start server at specified port
func (s *Server) Start(address string) error {
	return s.Echo.Start(address)
}
