package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Guaderxx/gowebtmpl/pkg/alog"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	Name string `json:"name"`
	ID   uint64 `json:"id"`
	jwt.RegisteredClaims
}

type CustomRefreshClaims struct {
	ID uint64 `json:"id"`
	jwt.RegisteredClaims
}

func JwtAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			if _, err := IsAuthorized(authToken, secret); err != nil {
				alog.Fatal("authorized failed", "error", err.Error())
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"code":  401,
					"error": "unauthorized",
				})
				return
			}
			userID, err := ExtractIDFromToken(authToken, secret)
			if err != nil {
				alog.Fatal("extract id failed", "error", err.Error())
				c.AbortWithStatusJSON(http.StatusOK, gin.H{
					"code":  401,
					"error": "unauthorized",
				})
				return
			}
			c.Set("x-user-id", userID)
			c.Next()
			return
		}
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code":  401,
			"error": "unauthorized",
		})
	}
}

func CreateAccessToken(id uint64, name, secret string, expiry int) (string, error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claims := &CustomClaims{
		Name: name,
		ID:   id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(exp, 0)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("signed token failed: %s", err)
	}
	return t, err
}

func CreateRefreshToken(id uint64, secret string, expiry int) (string, error) {
	claims := CustomRefreshClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("signed refresh token err: %s", err)
	}
	return rt, nil
}

func IsAuthorized(requestToken, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		alog.Error("authorized failed", "token", requestToken, "secret", secret, "error", err.Error())
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(refreshToken, secret string) (uint64, error) {
	token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return 0, fmt.Errorf("invalid token")
	} else {
		return claims["id"].(uint64), nil
	}
}
