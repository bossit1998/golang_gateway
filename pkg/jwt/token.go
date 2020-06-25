package jwt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//GenerateJWT - generates jwt jokens
func GenerateJWT(id, role string, signinigKey []byte) (access string, err error) {
	var (
		accessToken  *jwt.Token
		claims       jwt.MapClaims
	)
	accessToken = jwt.New(jwt.SigningMethodHS256)

	claims = accessToken.Claims.(jwt.MapClaims)
	claims["iss"] = "user"
	claims["sub"] = id
	claims["role"] = role
	claims["iat"] = time.Now().Unix()

	accessTokenString, err := accessToken.SignedString(signinigKey)
	if err != nil {
		err = fmt.Errorf("access_token generating error: %s", err)
		return
	}

	return accessTokenString, nil
}

func GenJWT(m map[interface{}]interface{}, key []byte) (access, refresh string, err error) {
	var (
		accessToken, refreshToken  *jwt.Token
		claims      jwt.MapClaims
	)
	accessToken = jwt.New(jwt.SigningMethodHS256)
	refreshToken = jwt.New(jwt.SigningMethodHS256)

	claims = accessToken.Claims.(jwt.MapClaims)
	rClaims := refreshToken.Claims.(jwt.MapClaims)

	for key, value := range m {
		claims[key.(string)] = value
		rClaims[key.(string)] = value
	}

	claims["iss"] = "user"
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().AddDate(0,0,3).Unix()

	rClaims["iss"] = "user"
	rClaims["iat"] = time.Now().Unix()
	rClaims["exp"] = time.Now().AddDate(0,0,15).Unix()

	accessTokenString, err := accessToken.SignedString(key)

	if err != nil {
		err = fmt.Errorf("access_token generating error: %s", err)
		return
	}

	refreshTokenString, err := refreshToken.SignedString(key)

	if err != nil {
		err = fmt.Errorf("refresh_token generating error: %s", err)
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
