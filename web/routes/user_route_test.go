package routes

import (
	"github.com/kataras/iris/v12/httptest"
	"testing"
)

func TestUserRoute(t *testing.T) {
	e := httptest.New(t, testApp())
	t.Log(e.GET("/users/uuid").Expect().Status(httptest.StatusBadRequest).Body())
}
