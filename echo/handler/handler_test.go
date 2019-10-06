package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	mockDB = map[string]*User{
		"jon@labstack.com": &User{"Jon Snow", "jon@labstack.com"},
		"bob@labstack.com": &User{"Bob Leaves", "bob@labstack.com"},
	}
	errorMockDB = map[string]*User{} 
	userJSON = `{"name":"Jon Snow","email":"jon@labstack.com"}`
	errorUserJSON = `{"msr:Jon Snow","rso;":"jon@labstack.com"}`
	usersJSON = `{{"name":"Jon Snow","email":"jon@labstack.com"},{"name":"Bob Leaves","email":"bob@labstack.com"}}`
)


/*
正常系のテスト
*/
func TestInsertHello(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/hey", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	cache = mockDB

	// Assertions
	if assert.NoError(t, InsertHello(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		// assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestGetHello(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/hey", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/hey/:email")
	c.SetParamNames("email")
	c.SetParamValues("jon@labstack.com")
	cache = mockDB

	// Assertions
	if assert.NoError(t, GetHello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// assert.Equal(t, userJSON, rec.Body.String())
	}
}

func TestHello(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	cache = mockDB

	// Assertions
	if assert.NoError(t, Hello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		// assert.Equal(t, userJSON, rec.Body.String())
	}
}


/*
異常系テスト
*/
func TestInsertHelloNoParams(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/hey", strings.NewReader(errorUserJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	cache = map[string]*User{} 

	// Assertions
	if assert.Error(t, InsertHello(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestGetHelloNoParams(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/hey", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/hey/:email")
	c.SetParamNames("email")
	c.SetParamValues("failed@email.com")
	cache = map[string]*User{} 

	// Assertions
	if assert.Error(t, GetHello(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}
