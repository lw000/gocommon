package tyutils

import (
	"errors"
	"reflect"
	"strconv"
)

func GetFloat64(v interface{}) (float64, error) {
	if v == nil {
		return 0, errors.New("v is nil")
	}

	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Float64 {
		return v.(float64), nil
	}
	return 0, errors.New("数据类型错误")
}

func GetInt(v interface{}) (int, error) {
	if v == nil {
		return 0, errors.New("v is nil")
	}
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Int {
		return v.(int), nil
	}
	return 0, errors.New("数据类型错误")
}

func GetUint32(v interface{}) (uint32, error) {
	if v == nil {
		return 0, errors.New("v is nil")
	}
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Uint32 {
		return v.(uint32), nil
	}
	return 0, errors.New("数据类型错误")
}

func GetInt32(v interface{}) (int32, error) {
	if v == nil {
		return 0, errors.New("v is nil")
	}
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Int32 {
		return v.(int32), nil
	}
	return 0, errors.New("数据类型错误")
}

func GetInt64(v interface{}) (int64, error) {
	if v == nil {
		return 0, errors.New("v is nil")
	}
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Int64 {
		return v.(int64), nil
	}
	return 0, errors.New("数据类型错误")
}

func GetString(v interface{}) (string, error) {
	if v == nil {
		return "", errors.New("v is nil")
	}
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.String {
		return v.(string), nil
	}
	return "", errors.New("数据类型错误")
}

func ToFloat64(s string) (float64, error) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, err
	}
	return v, nil
}

func ToInt64(s string) (int64, error) {
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return v, nil
}

func ToUint(s string) (uint64, error) {
	v, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}

	return v, nil
}

func ToUint32(s string) (uint32, error) {
	v, err := ToInt32(s)
	if err != nil {
		return 0, err
	}

	return uint32(v), nil
}

func ToInt32(s string) (int32, error) {
	v, err := ToInt(s)
	if err != nil {
		return 0, err
	}
	return int32(v), nil
}

func ToInt(s string) (int, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return v, nil
}

func ToBool(s string) (bool, error) {
	v, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}
	return v, nil
}
