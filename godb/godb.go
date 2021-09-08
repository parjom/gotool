package godb

import (
	"database/sql"
	"log"
	//	"runtime/debug"
	//	"github.com/parjom/gotool/goutil"
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
