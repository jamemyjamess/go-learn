package middlewareConfig

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"os"
	"strings"
)

func Init(e *echo.Echo) {
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Skipper: func(c echo.Context) bool {
			return strings.Contains(c.Path(), "swagger")
		},
		Level: 5,
	}))
	SetCORS(e)
	LogURL(e)
}

func SetCORS(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))
}

func LogURL(e *echo.Echo) {
	timeFormat := "time=${time_custom}"
	method := "method=${method}"
	url := "url=${host}${uri}"
	status := "http_status=${status}"
	errorLog := "error=${error}"
	remoteIP := "remote_ip=${remote_ip}"
	userAgent := "user_agent=${user_agent}"
	latencyHuman := "lantency=${latency_human}"
	apiLogFile, err := os.OpenFile(os.Getenv("STATIC_LOG_PATH")+"api-log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Println("Log File Error: ", err.Error())
	}

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           fmt.Sprintf("%s\n%s, %s, %s\n%s, %s, %s\n%s\n\n", timeFormat, method, status, url, remoteIP, latencyHuman, userAgent, errorLog),
		CustomTimeFormat: "2006/01/02 15:04:05",
	}))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           fmt.Sprintf("%s\n%s, %s, %s\n%s, %s, %s\n%s\n\n", timeFormat, method, status, url, remoteIP, latencyHuman, userAgent, errorLog),
		CustomTimeFormat: "2006/01/02 15:04:05",
		Output:           apiLogFile,
	}))
}
