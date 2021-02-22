package main

import (
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

func TestPostPunch(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app.Application)
	t.Log(e.POST("/punch/0").WithJSON(map[string]string{"monring": "1", "noon": "1", "night": "1"}).WithBasicAuth("feiyu", "11118888").Expect().Status(httptest.StatusOK).Body())
}
