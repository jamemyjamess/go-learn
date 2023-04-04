package keyConfig

import (
	"crypto/rsa"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/ssh"
)

var PublicKey *rsa.PublicKey
var PrivateKey interface{}

var PublicKeyProjectPublic *rsa.PublicKey

func Init() {
	LoadPublicKey()
	LoadPrivateKey()
	LoadPublicKeyProjectPublic()
}

func LoadPublicKey() {
	publicKeyPath, err := filepath.Abs(os.Getenv("STATIC_JWT_KYE_PATH") + "public.key")
	if err != nil {
		panic(err.Error())
	}

	publicKeyByte, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		panic(err.Error())
	}

	PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyByte)
	if err != nil {
		panic(err.Error())
	}
}

func LoadPrivateKey() {
	privateKeyPath, err := filepath.Abs(os.Getenv("STATIC_JWT_KYE_PATH") + "private.key")
	if err != nil {
		panic(err.Error())
	}

	privateKeyByte, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		panic(err.Error())
	}

	PrivateKey, err = ssh.ParseRawPrivateKey(privateKeyByte)
	if err != nil {
		panic(err.Error())
	}
}

func LoadPublicKeyProjectPublic() {
	publicKeyPath, err := filepath.Abs(os.Getenv("STATIC_JWT_KYE_PATH") + "public_project_public.key")
	if err != nil {
		panic(err.Error())
	}

	publicKeyByte, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		panic(err.Error())
	}

	PublicKeyProjectPublic, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyByte)
	if err != nil {
		panic(err.Error())
	}
}
