package utils

import (
	"github.com/gofrs/uuid"
)

func GenerateUUID() string {

	uuidString, _ := uuid.NewV4()

	return uuidString.String()
}

// func HandleGenerateId(c echo.Context) error {

// 	return responseConfig.Handler(c).Success(echo.Map{
// 		"car_id": GenerateUUID(),
// 	})

// }

func IsValidUUID(u string) bool {
	_, err := uuid.FromString(u)
	return err == nil
}
