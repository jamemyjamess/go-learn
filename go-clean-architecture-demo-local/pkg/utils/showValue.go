package utils

import (
	"fmt"
	"strconv"
)

func GetCo2OnSticker(co2Emission string) string {

	result := fmt.Sprintf("%.0f", RoundStringFloatDecimalPointsCustom(0, co2Emission))

	return result
}

func GetCombinedOnSticker(fuelConsumptionCombined string) string {

	result := fmt.Sprintf("%.1f", RoundStringFloatDecimalPointsCustom(1, fuelConsumptionCombined))

	return result
}

func GetFuelCombinedKmLOnSticker(fuelConsumptionCombined string) string {
	result := ""
	if value, err := strconv.ParseFloat(fuelConsumptionCombined, 64); err == nil {
		valueString := fmt.Sprintf("%g", value)
		result = fmt.Sprintf("%.1f", RoundStringFloatDecimalPointsCustom(1, valueString))
		if value != 0 {
			value = 100 / value
			valueString := fmt.Sprintf("%g", value)
			result = fmt.Sprintf("%.1f", RoundStringFloatDecimalPointsCustom(1, valueString))
		}
	} else {
		result = "0.0"
	}

	return result
}

func GetInCityOnSticker(fuelConsumptionUrban string) string {
	result := fmt.Sprintf("%.1f", RoundStringFloatDecimalPointsCustom(1, fuelConsumptionUrban))

	return result
}

func GetOutCityOnSticker(fuelConsumptionExtraUrban string) string {
	result := fmt.Sprintf("%.1f", RoundStringFloatDecimalPointsCustom(1, fuelConsumptionExtraUrban))

	return result
}

func GetDistanceOnSticker(drivingRange string) string {
	result := fmt.Sprintf("%.1f", RoundStringFloatDecimalPointsCustom(1, drivingRange))

	return result
}

func GetEnergyRateOnSticker(electricEnergyConsumption string) string {
	result := fmt.Sprintf("%.0f", RoundStringFloatDecimalPointsCustom(0, electricEnergyConsumption))

	return result
}

func GetFloatOnCo2AndSafetyDoc(inputValue string) string {
	inputFloat, _ := strconv.ParseFloat(inputValue, 64)

	result := fmt.Sprintf("%g", inputFloat)

	return result
}
