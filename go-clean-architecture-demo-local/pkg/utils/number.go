package utils

import (
	"log"
	"math"
	"strconv"
)

func RoundFloatDecimalPointsCustom(decimalPointAmount int, value float64) float64 {
	return math.Round(value*math.Pow(10, float64(decimalPointAmount))) / math.Pow(10, float64(decimalPointAmount))
}

func RoundStringFloatDecimalPointsCustom(decimalPointAmount int, value string) float64 {
	if value == "" {
		return 0
	}

	valueFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Println(err.Error())
		return 0
	}

	valueRound := math.Round(valueFloat*math.Pow(10, float64(decimalPointAmount))) / math.Pow(10, float64(decimalPointAmount))
	return valueRound
}
