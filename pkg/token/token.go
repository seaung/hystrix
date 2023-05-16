package token

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Configuration struct {
	key         string
	identityKey string
}

var (
	HeaderMissingError = errors.New("the length of the `Authorization` header is zero")
	config             = Configuration{"is token string is not none", "identityKey"}
	once               sync.Once
)

func Init(key, identityKey string) {
	once.Do(func() {
		if key != "" {
			config.key = key
		}
		if identityKey != "" {
			config.identityKey = identityKey
		}
	})
}

func ParseToken(tokenString string, key string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(key), nil
	})
	if err != nil {
		return "", err
	}

	var identityKey string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		identityKey = claims[config.identityKey].(string)
	}

	return identityKey, nil
}

func ParseRequest(c *gin.Context) (string, error) {
	header := c.Request.Header.Get("Authorization")

	if len(header) == 0 {
		return "", HeaderMissingError
	}

	var t string

	fmt.Sscanf(header, "Bearer %s", &t)

	return ParseToken(t, config.key)
}

func SignToken(identityKey string) (tokenString string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		config.identityKey: identityKey,
		"nbf":              time.Now().Unix(),
		"iat":              time.Now().Unix(),
		"exp":              time.Now().Add(100000 * time.Hour).Unix(),
	})

	tokenString, err = token.SignedString([]byte(config.key))

	return
}
