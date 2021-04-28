package jwt

import (
	"github.com/kataras/iris/v12/middleware/jwt"
	"time"
	"ucenter/s/config"
)

const (
	accessTokenMaxAge  = 10 * time.Minute
	refreshTokenMaxAge = time.Hour
)

type UserClaims struct {
	ID       string `json:"user_id"`
	Username string `json:"username"`
}

var (
	privateKey, publicKey = jwt.MustLoadRSA(config.C.GetFile("rsa/rsa_captcha_private_key.pem"),
		config.C.GetFile("rsa/rsa_captcha_public_key.pem"))
)

func NewTokenPair(uuid, account string) (accessToken, refreshToken string) {
	refreshClaims := jwt.Claims{Subject: uuid}

	accessClaims := UserClaims{
		ID:       uuid,
		Username: account,
	}

	signer := jwt.NewSigner(jwt.RS256, privateKey, accessTokenMaxAge)

	tokenPair, err := signer.NewTokenPair(accessClaims, refreshClaims, refreshTokenMaxAge)
	if err != nil {
		panic(err)
	}

	return string(tokenPair.AccessToken), string(tokenPair.RefreshToken)
}

func RefreshToken(refreshToken, uuid string) bool {
	verifier := jwt.NewVerifier(jwt.RS256, publicKey)
	_, err := verifier.VerifyToken([]byte(refreshToken), jwt.Expected{Subject: uuid})
	return err == nil
}
