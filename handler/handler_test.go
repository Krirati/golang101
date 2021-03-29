package handler_test

import (
	"golang101/mock"
	"golang101/model"
	"golang101/server"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestCallPathParamShouldBeReturnThatParam(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/nine", nil)
	rec := httptest.NewRecorder()
	// c := e.NewContext(req, rec)
	// c.SetPath("/:name")
	// c.SetParamNames("name")
	// c.SetParamValues("nine")
	// handler.ServeHomePath(c)
	uls := mock.UserLoginLogDB{}

	server := server.Server{
		Echo:           e,
		UserLoginLogDB: uls,
	}
	server.ServerRoute()
	server.ServerHTTP(rec, req)

	// Assertions
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "Hello :nine", rec.Body.String())

}

func TestCallPathJson_ShouldReturnJson(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/่json", nil)
	rec := httptest.NewRecorder()
	uls := mock.UserLoginLogDB{}

	server := server.Server{
		Echo:           e,
		UserLoginLogDB: uls,
	}
	server.ServerRoute()
	server.ServerHTTP(rec, req)

	expected := `{"Username": "krirati", "Point":10.2}`
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, expected, strings.TrimSpace(rec.Body.String()))
}

func TestCallReceivePostBody_ReturnBody(t *testing.T) {
	data := `{"Username":"nine","Password":"test"}`
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(data))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	// c := e.NewContext(req, rec)
	// c.SetPath("/login")
	// handler.ServeHomePathReceivePostBody(c)
	uls := mock.UserLoginLogDB{}
	ulm := model.UserLoginLog{Username: "nine", Password: "test"}

	uls.On("Create", &ulm).Return(nil)

	server := server.Server{
		Echo:           e,
		UserLoginLogDB: uls,
	}
	server.ServerRoute()
	server.ServerHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, data, strings.TrimSpace(rec.Body.String()))
}
