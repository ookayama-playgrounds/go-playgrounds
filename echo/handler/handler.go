package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

var cache = make(map[string]*User)

func InsertHello(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	cache[u.Email] = u
	return c.JSON(http.StatusCreated, u)
}

func GetHello(c echo.Context) error {
	email := c.Param("email")
	user, ok := cache[email]
	if !ok {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, user)
}

func Hello(c echo.Context) error {
	return c.JSON(http.StatusOK, cache)
}