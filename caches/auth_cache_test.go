package caches

import (
	"testing"
	_ "ucenter/s/tests/testsimple"
)

func TestCache(t *testing.T) {
	t.Log(NewAuthCache().Key("hello").Incr())
	t.Log(NewAuthCache().Key("hello").Incr())
	t.Log(NewAuthCache().Key("hello").Clear())
	t.Log(NewAuthCache().Key("hello").Incr())
	t.Log(NewAuthCache().Key("hello").Decr())
}
