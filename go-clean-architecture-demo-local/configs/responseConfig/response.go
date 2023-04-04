package responseConfig

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

type (
	Context interface {
		Success(data interface{}) error
		Created(data interface{}) error
		Accepted() error
		BadRequest(error interface{}) error
		Unauthorized(error ...string) error
		Forbidden() error
		NotFound(error ...string) error
		MethodNotAllowed() error
		InternalServerError() error
		InternalServerErrorSetError(error interface{}) error
		InvalidLogin() error
		InvalidToken() error
		ExpireToken() error
	}
	customContext struct {
		c echo.Context
	}
)

func Handler(c echo.Context) Context {
	return &customContext{c: c}
}

func Init(e *echo.Echo) {
	e.HTTPErrorHandler = CustomHTTPErrorHandler
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	if he, ok := err.(*echo.HTTPError); ok {
		if fmt.Sprint(he.Message) == "missing or malformed jwt" {
			Handler(c).BadRequest("missing jwt")
		} else if he.Internal != nil {
			log.Print(he.Internal.Error())
			if he.Internal.Error() == "forbidden" {
				Handler(c).Forbidden()
			} else if he.Internal.Error() == "token contains an invalid number of segments" {
				Handler(c).InvalidLogin()
			} else if he.Internal.Error() == "crypto/rsa: verification error" {
				Handler(c).InvalidToken()
			} else if strings.Contains(he.Internal.Error(), "token is expired") {
				Handler(c).ExpireToken()
			} else {
				Handler(c).InternalServerError()
			}
		} else {
			switch he.Code {
			case http.StatusBadRequest:
				Handler(c).BadRequest(he.Message)
			case http.StatusUnauthorized:
				Handler(c).Unauthorized(he.Internal.Error())
			case http.StatusForbidden:
				Handler(c).Forbidden()
			case http.StatusNotFound:
				Handler(c).NotFound()
			case http.StatusMethodNotAllowed:
				Handler(c).MethodNotAllowed()
			default:
				Handler(c).InternalServerError()
			}
		}
	} else {
		Handler(c).InternalServerErrorSetError(err)
	}
}

func (customContext *customContext) Success(data interface{}) error {
	return customContext.c.JSON(200, echo.Map{
		"success": true,
		"code":    200,
		"data":    data,
		"message": "Successfully",
	})
}

func (customContext *customContext) Created(data interface{}) error {
	return customContext.c.JSON(201, echo.Map{
		"success": true,
		"code":    201,
		"data":    data,
		"message": "Created",
	})
}

func (customContext *customContext) Accepted() error {
	return customContext.c.JSON(202, echo.Map{
		"success": true,
		"code":    202,
		"message": "Accepted",
	})
}

func (customContext *customContext) BadRequest(error interface{}) error {
	return customContext.c.JSON(400, echo.Map{
		"success": false,
		"code":    400,
		"message": "Bad Request",
		"error":   error,
	})
}

func (customContext *customContext) Unauthorized(error ...string) error {
	var errorMessage string

	if len(error) > 0 {
		errorMessage = error[0]
	}

	return customContext.c.JSON(401, echo.Map{
		"success": false,
		"code":    401,
		"message": "Unauthorized",
		"error":   errorMessage,
	})
}

func (customContext *customContext) Forbidden() error {
	return customContext.c.JSON(403, echo.Map{
		"success": false,
		"code":    403,
		"message": "Forbidden",
	})
}

func (customContext *customContext) NotFound(error ...string) error {
	var errorMessage string

	if os.Getenv("MODE") != "prod" {
		if len(error) > 0 {
			errorMessage = error[0]
		}
	}

	return customContext.c.JSON(400, echo.Map{
		"success": false,
		"code":    404,
		"message": "Not Found",
		"error":   errorMessage,
	})
}

func (customContext *customContext) MethodNotAllowed() error {
	return customContext.c.JSON(405, echo.Map{
		"success": false,
		"code":    405,
		"message": "Method Not Allowed",
	})
}

func (customContext *customContext) InternalServerError() error {
	return customContext.c.JSON(500, echo.Map{
		"success": false,
		"code":    500,
		"message": "Internal Server Error",
	})
}

func (customContext *customContext) InternalServerErrorSetError(err interface{}) error {
	return customContext.c.JSON(500, echo.Map{
		"success": false,
		"code":    500,
		"message": "Internal Server Error: " + err.(error).Error(),
		"error":   err,
	})
}

func (customContext *customContext) InvalidLogin() error {
	return customContext.c.JSON(401, echo.Map{
		"success": false,
		"code":    4010,
		"message": "Unauthorized",
		"error":   "invalid username or password",
	})
}

func (customContext *customContext) InvalidToken() error {
	return customContext.c.JSON(401, echo.Map{
		"success": false,
		"code":    4011,
		"message": "Unauthorized",
		"error":   "invalid token",
	})
}

func (customContext *customContext) ExpireToken() error {
	return customContext.c.JSON(401, echo.Map{
		"success": false,
		"code":    4012,
		"message": "Unauthorized",
		"error":   "token is expired",
	})
}
