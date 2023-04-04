package utils

import (
	"strconv"
	"time"
)

func MonthEngToThai(date time.Time) string {
	dateMonth := int(date.Month())
	switch dateMonth {
	case 1:
		return "มกราคม"
	case 2:
		return "กุมภาพันธ์"
	case 3:
		return "มีนาคม"
	case 4:
		return "เมษายน"
	case 5:
		return "พฤษภาคม"
	case 6:
		return "มิถุนายน"
	case 7:
		return "กรกฎาคม"
	case 8:
		return "สิงหาคม"
	case 9:
		return "กันยายน"
	case 10:
		return "ตุลาคม"
	case 11:
		return "พฤศจิกายน"
	case 12:
		return "ธันวาคม"
	default:
		return ""
	}
}

func YearEngToThai(date time.Time) string {
	year := date.Year()

	year += 543

	return strconv.Itoa(year)
}
