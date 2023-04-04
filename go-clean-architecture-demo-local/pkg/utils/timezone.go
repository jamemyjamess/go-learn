package utils

import (
	"reflect"
	"time"
)

func ConvertTimeZone(convertTime *time.Time) *time.Time {
	if convertTime != nil {
		var bangkokTimezone = *convertTime
		res := bangkokTimezone.Add(-time.Hour * 7)
		return &res
	}
	return convertTime
}

func TimeSubtract7HourOfStruct(modelStrcut interface{}) interface{} {

	v := reflect.ValueOf(modelStrcut).Elem()
	// log.Println(reflect.ValueOf(v))

	for i := 0; i < v.NumField(); i++ {
		switch timeInside := v.Field(i).Interface().(type) {
		case *time.Time:
			if timeInside == nil {
				continue
			}
			// log.Printf("t1: %T\n", h)
			var timeNew time.Time
			timeNew = *timeInside
			result := timeNew.Add(-time.Hour * 7)
			v.Field(i).Set(reflect.ValueOf(&result))
			// log.Println(v.Field(i))
		}
	}

	return &modelStrcut
}
