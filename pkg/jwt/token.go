package jwt

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"math/big"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/hkra/go-jwks"
)

//GenerateJWT - generates jwt jokens
func GenerateJWT(id, uType, role string, signinigKey []byte, clubId ...string) (access, refresh string, err error) {
	var (
		accessToken  *jwt.Token
		refreshToken *jwt.Token
		claims       jwt.MapClaims
	)
	accessToken = jwt.New(jwt.SigningMethodHS256)
	refreshToken = jwt.New(jwt.SigningMethodHS256)

	claims = accessToken.Claims.(jwt.MapClaims)
	claims["iss"] = "user"
	claims["sub"] = id
	claims["u_type"] = uType
	//@TODO should be fixed to minutes
	claims["exp"] = time.Now().Add(time.Hour * 500).Unix()
	claims["iat"] = time.Now().Unix()
	claims["role"] = role
	if len(clubId) == 1 {
		claims["club_id"] = clubId[0]
	}

	rClaims := refreshToken.Claims.(jwt.MapClaims)
	rClaims["iss"] = "user"
	rClaims["sub"] = id
	rClaims["iat"] = time.Now().Unix()

	accessTokenString, er := accessToken.SignedString(signinigKey)

	if er != nil {
		err = fmt.Errorf("access_token generating error: %s", er)
		return
	}

	refreshTokenString, er := refreshToken.SignedString(signinigKey)

	if er != nil {
		err = fmt.Errorf("access_token generating error: %s", er)
		return
	}

	return accessTokenString, refreshTokenString, nil
}

//ExtractClaims extracts claims from given token
func ExtractClaims(tokenStr string, signinigKey []byte) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err error
	)
	token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return signinigKey, nil
	})
	if err != nil {
		token, err = jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			// check token signing method etc
			return signinigKey, nil
		})
		if err != nil {
			return nil, err
		}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		err = fmt.Errorf("Invalid JWT Token")
		return nil, err
	}
	return claims, nil
}

//JWKSetToPublicKey ...
func JWKSetToPublicKey(kid string) ([]byte, error) {
	clientConfig := jwks.NewConfig()
	jwkClient := jwks.NewClient("https://appleid.apple.com/auth/keys", clientConfig)

	publicKey, err := jwkClient.GetSigningKey(kid)
	if err != nil {
		return nil, err
	}

	if publicKey == nil {
		fmt.Println("public key is nil")
	}

	jwk := map[string]string{}
	jwk["kty"] = publicKey.Kty
	jwk["n"] = publicKey.N
	jwk["e"] = publicKey.E

	if jwk["kty"] != "RSA" {
		err = fmt.Errorf("invalid key type: %s", jwk["kty"])
	}

	// decode the base64 bytes for n
	nb, err := base64.RawURLEncoding.DecodeString(jwk["n"])
	if err != nil {
		return nil, err
	}

	e := 0
	// The default exponent is usually 65537, so just compare the
	// base64 for [1,0,1] or [0,1,0,1]
	if jwk["e"] == "AQAB" || jwk["e"] == "AAEAAQ" {
		e = 65537
	} else {
		// need to decode "e" as a big-endian int
		err = fmt.Errorf("need to deocde e: %s", jwk["e"])
	}

	pk := &rsa.PublicKey{
		N: new(big.Int).SetBytes(nb),
		E: e,
	}

	der, err := x509.MarshalPKIXPublicKey(pk)
	if err != nil {
		return nil, err
	}

	block := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: der,
	}

	var out bytes.Buffer
	pem.Encode(&out, block)

	return out.Bytes(), nil
}
