package middlewareConfig

import (
	"errors"
	"fmt"
	"log"

	"github.com/jamemyjamess/go-clean-architecture-demo/configs/keyConfig"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Identity struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	RefId    string `json:"ref_id"`
	Role     string `json:"role"`
}

func JwtMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "RS256" {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
				}
				return keyConfig.PublicKey, nil
			}

			token, err := jwt.ParseWithClaims(auth, &jwt.StandardClaims{}, keyFunc)
			if err != nil {
				return nil, err
			}

			errValidateToken := errors.New("invalid or expired jwt")

			if token.Valid == false {
				return nil, errValidateToken
			}

			claims, ok := token.Claims.(*jwt.StandardClaims)

			if ok == false || token.Valid == false {
				return nil, errValidateToken
			}

			identity := new(Identity)
			identity.Id = claims.Subject
			// err = identity.Validate()

			if err != nil {
				log.Print(err.Error())
				return nil, err
			}

			sess, err := session.Get("session", c)

			if err != nil {
				return nil, err
			}

			sess.Values["identity"] = identity
			err = sess.Save(c.Request(), c.Response())

			if err != nil {
				return nil, err
			}

			return token, nil
		},
	})
}

func JwtWithSkipperMiddleware(skipper middleware.Skipper) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		TokenLookup:   "header:" + echo.HeaderAuthorization,
		SigningMethod: "RS256",
		AuthScheme:    "Bearer",
		SigningKey:    keyConfig.PublicKey,
		Skipper:       skipper,
	})
}

func JwtWithRoleMiddleware(role string) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "RS256" {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
				}
				return keyConfig.PublicKey, nil
			}

			token, err := jwt.ParseWithClaims(auth, &jwt.StandardClaims{}, keyFunc)
			if err != nil {
				return nil, err
			}

			errValidateToken := errors.New("invalid or expired jwt")

			if token.Valid == false {
				return nil, errValidateToken
			}

			claims, ok := token.Claims.(*jwt.StandardClaims)

			if ok == false || token.Valid == false {
				return nil, errValidateToken
			}

			identity := new(Identity)
			identity.Id = claims.Subject
			identity.Role = role
			// err = identity.ValidateRole()

			if err != nil {
				log.Print(err.Error())
				return nil, err
			}

			sess, err := session.Get("session", c)

			if err != nil {
				return nil, err
			}

			sess.Values["identity"] = identity
			err = sess.Save(c.Request(), c.Response())

			if err != nil {
				return nil, err
			}

			return token, nil
		},
	})
}
