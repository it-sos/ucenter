package jwt

import (
	"bytes"
	"testing"
	_ "ucenter/s/tests"
)

func TestNewTokenPair(t *testing.T) {
	NewTokenPair("uuid", "account")
	//t.Log(RefreshToken(string(tokenPair.RefreshToken[1:len(tokenPair.RefreshToken)-1]), "uuid", "account"))
	//tokenPair := NewTokenPair("uuid", "account")
	//t.Log(RefreshToken(tokenPair.RefreshToken, "uuid"))
	s := []byte(`eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJ1dWlkcyIsImlhdCI6MTYyMDc3NTI1MSwiZXhwIjoxNjIwNzc4ODUxfQ.OsviuOtFtvpH6smlwc1sjSVKsLkpmTp5dsUbypOslbrpuAfQmTso64sEJ8kGiNOj6xNHfHXRFfdHBgzR3Dy4FT-Oc2_gPM29ZGx4OEsHRpTYm-GJV5qDSwx_zmhiLKZRH6Neq6UMpp8Jhy00vOI45_XyZTe2bbEzVhF_acevo-4LimuVxspuPusUHrqr2upx5p1t7601gc9CUgQvFqaIsARDWuYqY-MB4jFhtQ-9OlHfyaq8vfMMtqn7rT42oOXzbFdwtRWGqVDiRcbV2QOIyI7DDzIe6t14KrclZ9gn0Wx-FGA0Y82tMNdFDFjfYDnI6YgOaJqiHDLQcYVr-gucqA`)
	//s := tokenPair.RefreshToken
	//t.Log(string(s))
	t.Log(ValidatorToken(bytes.Replace(s, []byte(`"`), []byte(``), 2)))
	//t.Log(string(tokenPair.RefreshToken))
	//t.Log(ValidatorToken(bytes.Replace(tokenPair.RefreshToken, []byte(`"`), []byte(``), 2)))
}

func TestRefreshToken(t *testing.T) {
}
