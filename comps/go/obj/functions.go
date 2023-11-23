package obj

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// ToInt convert interface to int32, if val is not a valid int32 value, zero will be returned
func ToUInt32(val interface{}) uint32 {
	var strValue string
	strValue = ToString("", val)

	floatValue, err := strconv.ParseUint(strValue, 10, 32)
	if err != nil {
		return uint32(0)
	}
	return uint32(floatValue)
}

// FromStringArray converte um array de string em um array de interfce
func FromStringArray(strArray []string) (ret []interface{}) {
	for _, value := range strArray {
		ret = append(ret, value)
	}
	return
}

// ToString converte uma interface em string
func ToString(defaultValue string, val interface{}) (strValue string) {
	switch val.(type) {
	case string:
		strValue = fmt.Sprintf("%s", val)
	case int64:
		strValue = fmt.Sprintf("%d", val)
	case *string:
		strValue = fmt.Sprintf("%s", *val.(*string))
	case int:
		strValue = fmt.Sprintf("%d", val)
	case *int:
		d := val.(*int)
		strValue = fmt.Sprintf("%d", *d)
	case *int64:
		d := val.(*int)
		strValue = fmt.Sprintf("%d", *d)
	case *int32:
		d := val.(*int)
		strValue = fmt.Sprintf("%d", *d)
	default:
		return defaultValue
	}
	if strValue == "" {
		return defaultValue
	}
	return strValue
}

func ToDatetime(val interface{}) (result time.Time) {
	if v, ok := val.(time.Time); ok {
		result = v
	}
	return result
}

func ToInt(val interface{}) (result int) {
	if v, ok := val.(int); ok {
		result = v
	}
	return result
}

func ToFloat(val interface{}) (result float64) {
	if v, ok := val.(float64); ok {
		result = v
	}
	return result
}

// ToIntArray convert interface into array of int
func ToIntArray(val interface{}) []int {
	var result []int
	if IsInterfaceNil(val) {
		return result
	}

	switch x := val.(type) {
	case []int:
		for _, i := range x {
			result = append(result, i)
		}
	case []interface{}:
		for _, i := range x {
			result = append(result, i.(int))
		}
	case []*int:
		for _, i := range x {
			result = append(result, *i)
		}
	default:
		fmt.Printf("Unsupported type: %T\n", x)
	}
	return result
}

// ToStringArray convert interface into array of string
func ToStringArray(val interface{}) (result []string) {
	if IsInterfaceNil(val) {
		return
	}

	switch x := val.(type) {
	case []interface{}:
		for _, v := range x {
			if value, ok := v.(string); ok {
				result = append(result, value)
			}
		}
	case []string:
		result = x
	}
	return
}

func IsInterfaceNil(value interface{}) bool {
	return value == nil || (reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil())
}

func IsTime(value interface{}) bool {
	_, ok := value.(time.Time)
	return ok
}

func ToBool(defaultValue bool, value interface{}) bool {
	v, ok := value.(bool)
	if !ok {
		return defaultValue
	}
	return v
}

func IsInterfacePointer(v interface{}) bool {
	return reflect.ValueOf(v).Type().Kind() == reflect.Ptr
}

func IsInstanceOf(objectPtr, typePtr interface{}) bool {
	return reflect.TypeOf(objectPtr) == reflect.TypeOf(typePtr)
}

func IsZeroValue(x interface{}) bool {
	return x == reflect.Zero(reflect.TypeOf(x)).Interface()
}

func GetFieldValue(v interface{}, fieldName string) (result string) {
	if field, ok := reflect.TypeOf(v).Elem().FieldByName(fieldName); ok {
		if field.Type.Kind() == reflect.String {
			result = field.Type.String()
		}
	}
	return
}

//
//func GetValueFromField(object interface{}, fieldName string) (any, error) {
//	objectValue := reflect.ValueOf(object)
//	if objectValue.Kind() == reflect.Ptr {
//		objectValue = objectValue.Elem()
//	}
//	if objectValue.Kind() != reflect.Struct {
//		return nil, ErrNonStructObject
//	}
//
//	for i := 0; i < objectValue.NumField(); i++ {
//		field := objectValue.Type().Field(i)
//		if strings.EqualFold(field.Name, fieldName) {
//			return objectValue.Field(i).Interface(), nil
//		}
//	}
//
//	return nil, ErrStructFieldNotFound
//}
