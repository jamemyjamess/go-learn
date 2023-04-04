package utils

import (
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateString() string {

	const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)

}

func GeneratePasswordHash(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return string(passwordHash), nil
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func StringToPrice(price string) string {
	result := ""
	p := message.NewPrinter(language.English)

	if price == "0" || price == "" {
		result = "-"
	} else {
		priceInt, _ := strconv.Atoi(SpaceStringsBuilder(price))
		result = p.Sprintf("%d", priceInt)
	}

	return result

}

func GetLinkToUTF8(link string) string {
	link = SpaceStringsBuilder(link)

	resultString := ""
	for _, charIndex := range link {
		char := fmt.Sprintf("%c", charIndex) //แปลงเป็นตัวหนังสือ
		value := ""
		if IsThaiString(char) {
			value = value + url.QueryEscape(char)
		} else if char == "/" {
			value = value + url.QueryEscape(char)
		} else {
			value = char
		}

		resultString = resultString + value
	}

	return resultString
}

//SpaceStringsBuilder ตัดช่องว่างออกทั้งหมด
func SpaceStringsBuilder(str string) string {
	var b strings.Builder
	b.Grow(len(str))
	for _, ch := range str {
		if !unicode.IsSpace(ch) {
			b.WriteRune(ch)
		}
	}
	return b.String()
}

func IsThaiString(input string) bool {

	var isThai = false

	for _, v := range input {
		if unicode.In(v, unicode.Thai) {
			isThai = true
		} else {
			isThai = false
		}
	}
	return isThai
}

func SetStringToDash(modelStrcut interface{}, useZeroValue bool) interface{} {

	v := reflect.ValueOf(modelStrcut).Elem()

	for i := 0; i < v.NumField(); i++ {
		switch value := v.Field(i).Interface().(type) {
		case string:
			if value == "" {
				v.Field(i).Set(reflect.ValueOf("-"))
				continue
			}
			if useZeroValue {
				if value == "0" {
					v.Field(i).Set(reflect.ValueOf("-"))
					continue
				}
			}
		}
		// if fmt.Sprintf("%v", v.Field(i).Kind()) == "struct" {

		// }
	}

	return &modelStrcut
}

func ToAlphaString(value int) string {
	if value < 0 {
		return ""
	}
	var ans string
	i := value
	for i > 0 {
		ans = string(rune(i-1)%26+65) + ans
		i = (i - 1) / 26
	}
	return ans
}

// func SetStrinngToDash2(modelStrcut interface{}) interface{} {

// 	log.Println(modelStrcut)

// 	return &modelStrcut
// }
