package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

func TestApp(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app.Application)
	t.Log(e.GET("/").WithBasicAuth(
		"feiyu",
		"11118888").Expect().Status(httptest.StatusOK).Body())
}

func TestDelete(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app.Application)

	r := e.DELETE("/users").WithJSON(iris.Map{"id": "1"}).Expect().
		Status(httptest.StatusBadRequest).
		Body()

	t.Log(r)
}
