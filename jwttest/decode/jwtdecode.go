package decode

import (
	"errors"
	"go-learn/jwttest/key"
	"net/http"
	"reflect"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AdminTokenClaim struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	Permission string `json:"permission"`
	Type       int    `json:"type"`
	Role       string `json:"role"`
	jwt.StandardClaims
}

type CompanyTokenClaim struct {
	ID              string `json:"id"`
	Username        string `json:"username"`
	CompanyFullName string `json:"company_full_name"`
	Type            int    `json:"type"`
	Role            string `json:"role"`
	jwt.StandardClaims
}

func ConfigTokenCustom() echo.MiddlewareFunc {
	_, publicKey := key.GetKey()
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract the token from Authorization header
			authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
			if authHeader == "" {
				return &echo.HTTPError{
					Code:     http.StatusUnauthorized,
					Message:  "Missing token",
					Internal: errors.New("Missing token"),
				}
			}
			parts := strings.SplitN(authHeader, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				return &echo.HTTPError{
					Code:     http.StatusUnauthorized,
					Message:  "Invalid token format",
					Internal: errors.New("Invalid token format"),
				}
			}

			// Define the claims model and the key to set in the context
			var claimsModel jwt.Claims
			var contextKey string

			// Check the role claim in the token to determine the claims model and the context key
			token, err := jwt.Parse(parts[1], func(token *jwt.Token) (interface{}, error) {
				// Check the signing method
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, errors.New("Unexpected signing method")
				}
				// Check the "role" claim
				if claims, ok := token.Claims.(jwt.MapClaims); ok && claims["role"] == "admin" {
					claimsModel = &AdminTokenClaim{}
					contextKey = "admin"
				} else if claims, ok := token.Claims.(jwt.MapClaims); ok && claims["role"] == "company" {

					claimsModel = &CompanyTokenClaim{}
					contextKey = "company"
				} else {
					return nil, errors.New("Invalid role claim")
				}
				return publicKey, nil
			})
			if err != nil && !token.Valid {
				return &echo.HTTPError{
					Code:     http.StatusUnauthorized,
					Message:  "Invalid or expired JWT",
					Internal: err,
				}
			}

			// Handle errors and set the claims in the context
			claims := reflect.New(reflect.ValueOf(claimsModel).Type().Elem()).Interface().(jwt.Claims)
			token, err = jwt.ParseWithClaims(parts[1], claims, func(token *jwt.Token) (interface{}, error) {
				// Check the signing method
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, errors.New("Unexpected signing method")
				}
				return publicKey, nil
			})
			if err != nil {
				return &echo.HTTPError{
					Code:     http.StatusUnauthorized,
					Message:  "Invalid or expired JWT",
					Internal: err,
				}
			}
			c.Set(contextKey, token)
			return next(c)
		}
	}
}
