package sessionConfig

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"os"
)

func Init(e *echo.Echo) {
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, _ := session.Get("session", c)
			sess.Options = &sessions.Options{
				Path:     "/",
				MaxAge:   0,
				HttpOnly: true,
				Secure:   true,
			}

			return next(c)
		}
	})
}
