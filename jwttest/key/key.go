package key

import (
	"crypto/rsa"
	"io/ioutil"
	"path/filepath"

	"github.com/golang-jwt/jwt"
)

func GetKey() (*rsa.PrivateKey, *rsa.PublicKey) {
	keyPrivatePath, _ := filepath.Abs("./key/private.key")
	key, err := ioutil.ReadFile(keyPrivatePath)
	if err != nil {
		panic(err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		panic(err)
	}
	keyPublicPath, _ := filepath.Abs("./keys/public.key")
	key, err = ioutil.ReadFile(keyPublicPath)
	if err != nil {
		panic(err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(key)
	if err != nil {
		panic(err)
	}
	return privateKey, publicKey
}

func d() {
	GetKey()

}
