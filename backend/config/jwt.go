package config

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	exp        string
	authorized bool
	email      string
}

func NewToken(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour * 200).Unix()
	claims["authorized"] = true
	claims["email"] = email

	private := PrivateKeyJwt()

	tokenString, err := token.SignedString([]byte(private))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenstring string) (bool, error) {

	token, err := jwt.Parse(tokenstring, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, fmt.Errorf("unauthorized")
		}

		private := PrivateKeyJwt()

		return []byte(private), nil
	})

	if token == nil {
		return false, nil
	}

	//parsear resultados
	if err != nil {
		return false, nil
	}

	//validar token
	if !token.Valid {
		return false, nil
	}
	return true, nil
}
