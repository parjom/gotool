package godb

import (
	"database/sql/driver"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func (db *DB) replaceParamQuery(query string, args Param) string {

	replKeyMap := make([]string,0)
	for key, value := range args {
		replKeyMap = append(replKeyMap, ":" + key, convertParams(value))
	}
    r := strings.NewReplacer(replKeyMap...)
	return r.Replace(query)
}

var convertibleTypes = []reflect.Type{reflect.TypeOf(time.Time{}), reflect.TypeOf(false), reflect.TypeOf([]byte{})}

func convertParams(v interface{}) string {
	const tmFmtWithMS = "2006-01-02 15:04:05.999"
	const tmFmtZero = "0000-00-00 00:00:00"
	const nullStr = "NULL"
	const escaper = "'"

	switch v := v.(type) {
	case RawValue:
		return string(v)
	case bool:
		return strconv.FormatBool(v)
	case time.Time:
		if v.IsZero() {
			return escaper + tmFmtZero + escaper
		} else {
			return escaper + v.Format(tmFmtWithMS) + escaper
		}
	case *time.Time:
		if v != nil {
			if v.IsZero() {
				return escaper + tmFmtZero + escaper
			} else {
				return escaper + v.Format(tmFmtWithMS) + escaper
			}
		} else {
			return nullStr
		}
	case driver.Valuer:
		reflectValue := reflect.ValueOf(v)
		if v != nil && reflectValue.IsValid() && ((reflectValue.Kind() == reflect.Ptr && !reflectValue.IsNil()) || reflectValue.Kind() != reflect.Ptr) {
			r, _ := v.Value()
			return convertParams(r)
		} else {
			return nullStr
		}
	case fmt.Stringer:
		reflectValue := reflect.ValueOf(v)
		if v != nil && reflectValue.IsValid() && ((reflectValue.Kind() == reflect.Ptr && !reflectValue.IsNil()) || reflectValue.Kind() != reflect.Ptr) {
			return escaper + strings.Replace(fmt.Sprintf("%v", v), escaper, "\\"+escaper, -1) + escaper
		} else {
			return nullStr
		}
	case []byte:
		if isPrintable(v) {
			return escaper + strings.Replace(string(v), escaper, "\\"+escaper, -1) + escaper
		} else {
			return escaper + "<binary>" + escaper
		}
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return toString(v)
	case float64, float32:
		return fmt.Sprintf("%.6f", v)
	case string:
		return escaper + strings.Replace(v, escaper, "\\"+escaper, -1) + escaper
	default:
		rv := reflect.ValueOf(v)
		if v == nil || !rv.IsValid() || rv.Kind() == reflect.Ptr && rv.IsNil() {
			return nullStr
		} else if valuer, ok := v.(driver.Valuer); ok {
			v, _ = valuer.Value()
			return convertParams(v)
		} else if rv.Kind() == reflect.Ptr && !rv.IsZero() {
			return convertParams(reflect.Indirect(rv).Interface())
		} else {
			for _, t := range convertibleTypes {
				if rv.Type().ConvertibleTo(t) {
					return convertParams(rv.Convert(t).Interface())
				}
			}
			return escaper + strings.Replace(fmt.Sprint(v), escaper, "\\"+escaper, -1) + escaper
		}
	}
}

func isPrintable(s []byte) bool {
	for _, r := range s {
		if !unicode.IsPrint(rune(r)) {
			return false
		}
	}
	return true
}

func mergeMaps(maps ...Param) Param {
    result := make(Param)
    for _, m := range maps {
        for k, v := range m {
            result[k] = v
        }
    }
    return result
}

func toString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.FormatInt(int64(v), 10)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	}
	return ""
}