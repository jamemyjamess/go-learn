package utils

import "strings"

var TypeWordStickerICE = "ICE"
var TypeWordStickerBEV = "BEV"
var TypeWordStickerPHEV = "PHEV"
var TyreWordStickerE85 = "E85"
var TypeWordStickerB10 = "B10"

func GetTypeSticker(carTypeOptionID, engineOptionID string) string {
	typeSticker := TypeWordStickerICE
	switch carTypeOptionID {
	case "0000000014":
		typeSticker = TypeWordStickerBEV
	case "0000000012", "0000000013":
		if engineOptionID == "e9071d7f06c84247869f4a12913c1ec0" {
			typeSticker = TypeWordStickerPHEV
		}
	default:
		typeSticker = TypeWordStickerICE
	}

	return typeSticker

}

func GetSubTypeSticker(carTypeOptionID, engineOptionID string, fuelOptionID string) string {
	typeSticker := TypeWordStickerICE
	switch carTypeOptionID {
	case "0000000014":
		typeSticker = TypeWordStickerBEV
	case "0000000012", "0000000013":
		if engineOptionID == "e9071d7f06c84247869f4a12913c1ec0" {
			typeSticker = TypeWordStickerPHEV
		}
	case "0000000020", "0000000022", "0000000023":
		typeSticker = TyreWordStickerE85
	case "0000000021":
		if engineOptionID == "386ed8d2310b4eae8a2b48c191b6a586" && fuelOptionID == "634322ac22b345ea906c005285538ef7" {
			typeSticker = TypeWordStickerB10
		}
	default:
		typeSticker = TypeWordStickerICE
	}

	return typeSticker

}

func TireOnSticker(value string) string {
	if value == "" {
		return value
	}
	result := ""
	value = strings.ReplaceAll(value, "-", "/")
	sp := strings.Split(value, "/")
	for i, item := range sp {
		if i == 0 {
			result = item + "/"
			continue
		}
		result += item
	}

	return result
}

func TireOnStickerAppend(wheelSize, frontWheelSize, backWheelSize string, isParentheses bool) string {
	if wheelSize == "" && frontWheelSize == "" && backWheelSize == "" {
		return ""
	}

	result := ""
	result = wheelSize
	if result != "" {
		result = TireOnSticker(result)
	}

	resultfrontBackTire := ""
	if frontWheelSize != "" {
		frontWheelSize = TireOnSticker(frontWheelSize)
		resultfrontBackTire = frontWheelSize
	}

	if resultfrontBackTire != "" {
		backWheelSize = TireOnSticker(backWheelSize)
		resultfrontBackTire += ", " + backWheelSize
	} else {
		resultfrontBackTire = backWheelSize
	}
	if resultfrontBackTire != "" {
		if isParentheses {
			result = " ( " + resultfrontBackTire + " )"
		} else {
			result = resultfrontBackTire
		}
	}

	return result
}
