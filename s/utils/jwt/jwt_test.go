package jwt

import (
	"testing"
	_ "ucenter/s/tests"
)

func TestNewTokenPair(t *testing.T) {
	NewTokenPair("uuid", "account")
	//t.Log(RefreshToken(string(tokenPair.RefreshToken[1:len(tokenPair.RefreshToken)-1]), "uuid", "account"))
}

func TestRefreshToken(t *testing.T) {
}
