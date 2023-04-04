package utils

import (
	"encoding/json"
	"reflect"
)

// Set Difference: A - B
func DifferenceArrayString(a, b []string) (diff []string) {
	m := make(map[string]bool)

	for _, item := range b {
		m[item] = true
	}

	for _, item := range a {
		if _, ok := m[item]; !ok {
			diff = append(diff, item)
		}
	}
	return
}

func ChunksBySizeArrayString(xs []string, chunkSize int) [][]string {
	if len(xs) == 0 {
		return nil
	}
	divided := make([][]string, (len(xs)+chunkSize-1)/chunkSize)
	prev := 0
	i := 0
	till := len(xs) - chunkSize
	for prev < till {
		next := prev + chunkSize
		divided[i] = xs[prev:next]
		prev = next
		i++
	}
	divided[i] = xs[prev:]
	return divided
}

func StructToResponse(structModel interface{}) map[string]interface{} {
	newPayload := NilSliceToEmptySlice(structModel)

	var result map[string]interface{}
	byteData, _ := json.Marshal(newPayload)
	json.Unmarshal(byteData, &result)

	return result
}

func NilSliceToEmptySlice(inter interface{}) interface{} {
	// original input that can't be modified
	val := reflect.ValueOf(inter)

	switch val.Kind() {
	case reflect.Slice:
		newSlice := reflect.MakeSlice(val.Type(), 0, val.Len())
		if !val.IsZero() {
			// iterate over each element in slice
			for j := 0; j < val.Len(); j++ {
				item := val.Index(j)

				var newItem reflect.Value
				switch item.Kind() {
				case reflect.Struct:
					// recursively handle nested struct
					newItem = reflect.Indirect(reflect.ValueOf(NilSliceToEmptySlice(item.Interface())))
				default:
					newItem = item
				}

				newSlice = reflect.Append(newSlice, newItem)
			}

		}
		return newSlice.Interface()
	case reflect.Struct:
		// new struct that will be returned
		newStruct := reflect.New(reflect.TypeOf(inter))
		newVal := newStruct.Elem()
		// iterate over input's fields
		for i := 0; i < val.NumField(); i++ {
			newValField := newVal.Field(i)
			valField := val.Field(i)
			switch valField.Kind() {
			case reflect.Slice:
				// recursively handle nested slice
				newValField.Set(reflect.Indirect(reflect.ValueOf(NilSliceToEmptySlice(valField.Interface()))))
			case reflect.Struct:
				// recursively handle nested struct
				newValField.Set(reflect.Indirect(reflect.ValueOf(NilSliceToEmptySlice(valField.Interface()))))
			default:
				newValField.Set(valField)
			}
		}

		return newStruct.Interface()
	case reflect.Map:
		// new map to be returned
		newMap := reflect.MakeMap(reflect.TypeOf(inter))
		// iterate over every key value pair in input map
		iter := val.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()
			// recursively handle nested value
			newV := reflect.Indirect(reflect.ValueOf(NilSliceToEmptySlice(v.Interface())))
			newMap.SetMapIndex(k, newV)
		}
		return newMap.Interface()
	case reflect.Ptr:
		// dereference pointer
		return NilSliceToEmptySlice(val.Elem().Interface())
	default:
		return inter
	}
}

func StructToResponseWithNilSliceToEmptySlice(structModel interface{}) interface{} {
	return NilSliceToEmptySlice(structModel)
}

func RemoveDuplicateSliceString(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
