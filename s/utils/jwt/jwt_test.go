package jwt

import (
	"testing"
	_ "ucenter/s/tests"
)

func TestNewTokenPair(t *testing.T) {
	_, refreshToken := NewTokenPair("uuid", "account")
	t.Log(refreshToken)
	t.Log(RefreshToken(refreshToken, "uuid"))
}

func TestRefreshToken(t *testing.T) {
}
