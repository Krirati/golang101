package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type ModelTest struct {
	Username string      `json: "username"`
	Password string      `json: "password"`
	Correct  interface{} `json: "correct"`
}

func GetImage(c echo.Context) error {
	c.Response().Header().Set("Accept-Ranges", "bytes")
	c.Response().Header().Set("Content-Type", "image/jpeg")
	c.Response().Header().Set("Connection", "keep-alive")
	c.Response().Header().Set("Cache-Control", "must-revalidate")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Cache-Control", "no-store")
	c.Response().Status = 206
	return c.File("D:/Downloads/image.jpg")
}
func GetJSON(c echo.Context) error {
	// c.Response().Header().Set("Accept-Ranges", "bytes")
	// c.Response().Header().Set("Content-Type", "application/json")
	// c.Response().Header().Set("Connection", "keep-alive")
	// c.Response().Header().Set("Cache-Control", "must-revalidate")
	// c.Response().Header().Set("Cache-Control", "no-cache")
	// c.Response().Header().Set("Cache-Control", "no-store")
	// c.Response().Status = 206
	return c.File("test.json")
}

type CatlogNodes struct {
	CatlogNodes []Catlog `json:"catlog_nodes"`
}
type Catlog struct {
	Product_id string `json: "product_id"`
	Quantity   int    `json: "quantity"`
}

func GetResponseJson(c echo.Context) error {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	rand.Seed(time.Now().UnixNano())

	for i := len(a) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)
	// result := GetJSON(c)
	file, _ := ioutil.ReadFile("D:/Nine/INTERN/file-face-labelling/gen/5089/label.json")
	data := CatlogNodes{}

	_ = json.Unmarshal([]byte(file), &data)
	// u := ModelTest{
	// 	Username: "nine",
	// 	Password: "1234",
	// 	// Correct:  result,
	// }
	return c.JSON(http.StatusOK, data)
}
