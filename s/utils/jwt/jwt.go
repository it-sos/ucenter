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
	ID      string `json:"uuid"`
	Account string `json:"account"`
}

var (
	privateKey, publicKey = jwt.MustLoadRSA(config.C.GetFile("rsa/rsa_captcha_private_key.pem"),
		config.C.GetFile("rsa/rsa_captcha_public_key.pem"))
)

func NewTokenPair(uuid, account string) jwt.TokenPair {
	refreshClaims := jwt.Claims{Subject: uuid}

	accessClaims := UserClaims{
		ID:      uuid,
		Account: account,
	}

	signer := jwt.NewSigner(jwt.RS256, privateKey, accessTokenMaxAge)

	tokenPair, err := signer.NewTokenPair(accessClaims, refreshClaims, refreshTokenMaxAge)
	if err != nil {
		panic(err)
	}

	return tokenPair
}

//func RefreshToken(refreshToken, uuid string) (jwt.TokenPair, error) {
//	verifier := jwt.NewVerifier(jwt.RS256, publicKey)
//	verifiedToken, err := verifier.VerifyToken([]byte(refreshToken), jwt.Expected{Subject: uuid})
//	if err != nil {
//		return jwt.TokenPair{}, err
//	}
//	fmt.Println(verifiedToken.StandardClaims.Subject)
//	fmt.Println(verifiedToken.StandardClaims)
//	fmt.Println(string(verifiedToken.Token))
//	fmt.Println(string(verifiedToken.Header))
//	fmt.Println(string(verifiedToken.Payload))
//}
//
//func ValidateToken(accessToken string) error {
//	verifier := jwt.NewVerifier(jwt.RS256, publicKey)
//	verifiedToken, err := verifier.VerifyToken([]byte(accessToken))
//	if err != nil {
//		return jwt.TokenPair{}, err
//	}
//}
