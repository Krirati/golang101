package main

import (
	"fmt"
	"golang101/model"
	"golang101/server"

	"github.com/labstack/echo"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	if db != nil {
		fmt.Println("connect to database")
	}
	// Migrate the schema
	db.AutoMigrate(&model.Product{}, &model.UserLoginLog{})

	e := echo.New()

	server := server.Server{
		Echo:           e,
		DB:             db,
		UserLoginLogDB: model.NewUserLoginLogDB(db),
	}

	server.ServerRoute()
	server.Start(":1323")
	// e.Logger.Fatal(e.Start(":1323"))
}
