package configs

import (
	"os"

	"github.com/jamemyjamess/go-clean-architecture-demo/configs/dbConfig"
	"github.com/jamemyjamess/go-clean-architecture-demo/configs/envConfig"
	"github.com/jamemyjamess/go-clean-architecture-demo/configs/keyConfig"
	"github.com/jamemyjamess/go-clean-architecture-demo/configs/logConfig"
	"github.com/jamemyjamess/go-clean-architecture-demo/configs/middlewareConfig"
	"github.com/jamemyjamess/go-clean-architecture-demo/configs/responseConfig"
	"github.com/jamemyjamess/go-clean-architecture-demo/configs/validatorConfig"

	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	logConfig.Init()
	envConfig.Init()
	CreateDirectory()
	responseConfig.Init(e)
	dbConfig.Init()
	middlewareConfig.Init(e)
	validatorConfig.Init(e)
	keyConfig.Init()
	// sessionConfig.Init(e)
}

func CreateDirectory() {
	var err error

	if err = os.MkdirAll(os.Getenv("STATIC_JWT_KYE_PATH"), os.ModePerm); err != nil {
		panic(err.Error())
	}

	if err = os.MkdirAll(os.Getenv("STATIC_LOG_PATH"), os.ModePerm); err != nil {
		panic(err.Error())
	}

	if err = os.MkdirAll(os.Getenv("STATIC_FILE_PATH"), os.ModePerm); err != nil {
		panic(err.Error())
	}

}
