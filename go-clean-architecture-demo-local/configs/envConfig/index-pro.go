//go:build pro

package envConfig

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

func Init() {
	LoadEnv()
}

var modeEnv = "pro"

func LoadEnv() {
	log.Println("modeEnv: ", modeEnv)
	path, _ := filepath.Abs("./.env.production")
	err := godotenv.Load(path)

	if err != nil {
		panic(err.Error())
	}
}

func GetEnvStr(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return v, errors.New("env variable empty")
	}
	return v, nil
}

func GetEnvInt(key string) (int, error) {
	s, err := GetEnvStr(key)
	if err != nil {
		return 0, err
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func GetEnvBool(key string) (bool, error) {
	s, err := GetEnvStr(key)
	if err != nil {
		return false, err
	}
	v, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}
	return v, nil
}

func GetEnvInt64(key string) (int64, error) {
	s, err := GetEnvStr(key)
	if err != nil {
		return 0, err
	}

	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return v, nil
}
