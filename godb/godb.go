package godb

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"
	//	"runtime/debug"
	"github.com/parjom/gotool/goutil"
)

type DB struct {
	Error  error
	db     *sql.DB
	debug  bool
	logger *log.Logger
}

type Param map[string]interface{}

func Open(driverName, dataSourceName string) (*DB, error) {
	db := &DB{
		Error:  nil,
		db:     nil,
		debug:  false,
		logger: log.Default(),
	}

	_db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	db.db = _db
	db.Error = err
	_db.Driver()
	return db, err
}
func (db *DB) SetDebug(debug bool) {
	db.debug = debug
}
func (db *DB) GetDebug() bool {
	return db.debug
}

func (db *DB) Exec(query string, args ...Param) (int64, int64, error) {
	if len(args) == 0 {
		// 파라메터가 없는 단순쿼리
		return db.execSimple(query)
	} else {
		param := mergeMaps(args...)
		return db.execParam(query, param)
	}
}

func (db *DB) execSimple(query string) (RowsAffected int64, LastInsertId int64, err error) {
	RowsAffected = -1
	LastInsertId = -1

	res, err := db.db.Exec(query)
	if err == nil {
		if ra, e := res.RowsAffected(); e == nil {
			RowsAffected = ra
		}
		if lii, e := res.LastInsertId(); e == nil {
			LastInsertId = lii
		}
	}
	return RowsAffected, LastInsertId, err
}

func (db *DB) execParam(query string, args Param) (int64, int64, error) {
	newQuery := db.replaceParamQuery(query, args)
	return db.execSimple(newQuery)
}

func (db *DB) replaceParamQuery(query string, args Param) string {

	// for key, value := range args {

	//     switch v := value.(type) {
	//     case int:
	//         // v is an int here, so e.g. v + 1 is possible.
	//         fmt.Printf("Integer: %v", v)
	//     case float64:
	//         // v is a float64 here, so e.g. v + 1.0 is possible.
	//         fmt.Printf("Float64: %v", v)
	//     case string:
	//         // v is a string here, so e.g. v + " Yeah!" is possible.
	//         fmt.Printf("String: %v", v)
	//     default:
	//         // And here I'm feeling dumb. ;)
	//         fmt.Printf("I don't know, ask stackoverflow.")
	//     }

	//     replacer := strings.NewReplacer("/v1.0/", "", "/emp_1/", "_")

	// }
	return ""
}

var convertibleTypes = []reflect.Type{reflect.TypeOf(time.Time{}), reflect.TypeOf(false), reflect.TypeOf([]byte{})}

func convertParams(v interface{}) string {
	const tmFmtWithMS = "2006-01-02 15:04:05.999"
	const tmFmtZero = "0000-00-00 00:00:00"
	const nullStr = "NULL"
	const escaper = "'"

	switch v := v.(type) {
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
		return goutil.ToString(v)
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
			vars[idx] = escaper + strings.Replace(fmt.Sprint(v), escaper, "\\"+escaper, -1) + escaper
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

