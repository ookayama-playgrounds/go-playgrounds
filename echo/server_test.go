package main

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/labstack/echo"
)

type testCase struct {
	Name     string
	Routes   []*echo.Route
	Handlers map[string]echo.HandlerFunc
}

func TestMain_Main(t *testing.T) {
	cases := []testCase{
		testCase{
			"none test",
			[]*echo.Route{},
			make(map[string]echo.HandlerFunc),
		},
		testCase{
			"single route test",
			[]*echo.Route{
				&echo.Route{http.MethodGet, "/hello", "sample get"},
			},
			map[string]echo.HandlerFunc{
				"get": func(c echo.Context) error { return c.String(http.StatusOK, "get ok") },
			},
		},
		testCase{
			"rest api route test",
			[]*echo.Route{
				&echo.Route{http.MethodGet, "/rest", "rest get"},
				&echo.Route{http.MethodPost, "/rest", "rest post"},
				&echo.Route{http.MethodPut, "/rest", "rest put"},
				&echo.Route{http.MethodPatch, "/rest", "rest patch"},
				&echo.Route{http.MethodDelete, "/rest", "rest delete"},
			},
			map[string]echo.HandlerFunc{
				"get":    func(c echo.Context) error { return c.String(http.StatusOK, "rest get ok") },
				"post":   func(c echo.Context) error { return c.String(http.StatusCreated, "rest post ok") },
				"put":    func(c echo.Context) error { return c.String(http.StatusNoContent, "rest put ok") },
				"patch":  func(c echo.Context) error { return c.String(http.StatusNoContent, "rest patch ok") },
				"delete": func(c echo.Context) error { return c.String(http.StatusNoContent, "rest delete ok") },
			},
		},
	}

	for _, c := range cases {
		fmt.Printf("Test case: %s\n", c.Name)
		e := echo.New()
		for _, route := range c.Routes {
			switch route.Method {
			case http.MethodGet:
				e.GET(route.Path, c.Handlers["get"]).Name = route.Name
			case http.MethodPost:
				e.POST(route.Path, c.Handlers["post"]).Name = route.Name
			case http.MethodPut:
				e.PUT(route.Path, c.Handlers["put"]).Name = route.Name
			case http.MethodPatch:
				e.PATCH(route.Path, c.Handlers["patch"]).Name = route.Name
			case http.MethodDelete:
				e.DELETE(route.Path, c.Handlers["delete"]).Name = route.Name
			default:
				t.Errorf("Invalid test case: %v\n", c)
			}
		}
		got := e.Routes()
		wanted := c.Routes
		if !reflect.DeepEqual(got, wanted) {
			t.Errorf("\n  wanted: [%v],\n  got: [%v]", wanted, got)
		}
	}
}
