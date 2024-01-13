package helper

import "reflect"

func IsEmptyStruct(value interface{}) bool {
	val := reflect.ValueOf(value)
	if val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			zeroValue := reflect.Zero(field.Type()).Interface()

			if !reflect.DeepEqual(field.Interface(), zeroValue) {
				return false
			}
		}
		return true
	}
	return false
}

func SetValue(value interface{}, defaultValue interface{}) interface{} {
	switch v := value.(type) {
	case string:
		if v != "" {
			return v
		}
	case int:
	case bool:
		return v
	case struct{}:
		if !IsEmptyStruct(v) {
			return v
		}
	}

	return defaultValue
}
