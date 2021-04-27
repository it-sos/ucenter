package jwt

import (
	"crypto/rsa"
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

func NewTokenPair(uuid, account string) jwt.TokenPair {
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

	return tokenPair
}

func RefreshToken(refreshToken, uuid string) {
	verifier := jwt.NewVerifier(jwt.RS256, publicKey)
	verifier.VerifyToken([]byte(refreshToken), jwt.Expected{Subject: uuid})
}

func getRsaPem() (*rsa.PrivateKey, *rsa.PublicKey) {
	return jwt.MustLoadRSA(config.C.GetFile("rsa/rsa_captcha_private_key.pem"),
		config.C.GetFile("rsa/rsa_captcha_public_key.pem"))
}
