package utils

import (
	"time"
)

func DepartmentDelayCondition(department, fastTyre, carTypeOptionID string, start, end *time.Time) (bool, int) {

	workday := getWorkday(start, end)
	var delay bool = false
	var delayDay int = 0

	//  limit day of oie, tisi, technical service
	var limitDay int = 7
	var exciseLimitDay int = 3

	if department == "oie" {
		if fastTyre == "tyre" {
			delay = false
			delayDay = 0
			return delay, delayDay
		}
	} else if department == "tisi" {
		if carTypeOptionID == "0000000014" {
			delay = false
			delayDay = 0
			return delay, delayDay
		}
	} else if department == "excise" {
		if workday > exciseLimitDay {
			delay = true
			delayDay = workday - exciseLimitDay
			return delay, delayDay
		}
	}

	if workday > limitDay {
		delay = true
		delayDay = workday - limitDay
	}

	return delay, delayDay
}

func GetApprovalAtLatest(OieApprovalAt, TisiApprovalAt, TsApprovalAt *time.Time) *time.Time {
	// หา date last
	var dateLast time.Time
	if dateLast.Before(*OieApprovalAt) {
		dateLast = *OieApprovalAt
	}

	if dateLast.Before(*TisiApprovalAt) {
		dateLast = *TisiApprovalAt
	}

	if dateLast.Before(*TsApprovalAt) {
		dateLast = *TsApprovalAt
	}

	return &dateLast
}

func GetApprovalAtLatest2(setTime ...*time.Time) *time.Time {
	// หา date last
	var dateLast time.Time
	for _, a := range setTime {
		if a != nil {
			if dateLast.Before(*a) {
				dateLast = *a
			}
		}
	}

	return &dateLast
}

func getWorkday(start *time.Time, end *time.Time) int {
	startAdd1 := start.AddDate(0, 0, 1)
	// var oieLimitDate = 20
	now := time.Now()
	if end == nil {
		end = &now
	}

	listDate := setArrayDate(startAdd1, *end)
	workday := getCountWorkday(listDate)

	// เหลือการลบวันหยุด

	return workday
}

func setArrayDate(startDate, endDate time.Time) []time.Time {
	var listDate []time.Time
	for rd := rangeDate(startDate, endDate); ; {
		date := rd()
		if date.IsZero() {
			break
		}
		listDate = append(listDate, date)
	}
	return listDate
}

func getCountWorkday(listDate []time.Time) int {
	var workday int = 0
	for _, date := range listDate {
		if isNotWeekend(date) {
			workday++
		}
	}
	return workday
}

func rangeDate(start, end time.Time) func() time.Time {
	y, m, d := start.Date()
	start = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	y, m, d = end.Date()
	end = time.Date(y, m, d, 0, 0, 0, 0, time.UTC)

	return func() time.Time {
		if start.After(end) {
			return time.Time{}
		}
		date := start
		start = start.AddDate(0, 0, 1)
		return date
	}
}

// A weekend is Friday 10pm UTC to Sunday 10:05pm UTC
func isNotWeekend(t time.Time) bool {
	t = t.UTC()

	switch t.Weekday() {
	case time.Saturday:
		return false
	case time.Sunday:
		h, m, _ := t.Clock()
		if h < 12+10 {
			return false
		}
		if h == 12+10 && m <= 5 {
			return false
		}
	}
	return true
}
