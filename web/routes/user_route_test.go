package routes

import (
	"github.com/kataras/iris/v12/httptest"
	"testing"
	"ucenter/s/tests/testapp"
)

func TestUserRoute(t *testing.T) {
	e := httptest.New(t, testapp.App(Configure))
	t.Log(e.GET("/users").Expect().Status(httptest.StatusBadRequest).Body())
}
