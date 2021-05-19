package main

import (
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

func TestApp(t *testing.T) {
	app := newApp()
	e := httptest.New(t, app.Application)
	t.Log(e.GET("/").Expect().Status(httptest.StatusBadRequest).Body().Equal("hello is error."))
}
