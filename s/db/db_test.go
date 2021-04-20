package db

import (
	"golang.org/x/net/context"
	"testing"
	_ "ucenter/s/tests"
)

func TestDds(t *testing.T) {
	Init()
	var ctx = context.Background()
	t.Log(Rdb.Set(ctx, "key", "value", 0).Err())
	t.Log(Rdb.Get(ctx, "key"))
}

func TestDd(t *testing.T) {
	Init()
	Conn.ShowExecTime(true)
}
