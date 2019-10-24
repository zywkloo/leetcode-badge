package tools

import (
	"reflect"
	"time"
)

// ZeroTime 返回当天的零点时间，时区 UTC+8
func ZeroTime(t time.Time) time.Time {

	t = t.Add(-1 * time.Duration(t.Hour()) * time.Hour)
	t = t.Add(-1 * time.Duration(t.Minute()) * time.Minute)
	t = t.Add(-1 * time.Duration(t.Second()) * time.Second)
	return t
}

// IsToday 判断时间是否是今天
func IsToday(t time.Time) bool {
	t1 := ZeroTime(t)
	t2 := ZeroTime(time.Now())

	return t1.Unix() == t2.Unix()
}

// BoolToInt bool to int
func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// IntToBool int to bool
func IntToBool(n int) bool {
	return n == 1
}

// Struct2Map struct to map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		tag, ok := t.Field(i).Tag.Lookup("json")
		if !ok {
			continue
		}
		data[tag] = v.Field(i).Interface()
	}
	return data
}
