package godb

import (
	"database/sql"
	"log"
)

type DB struct {
	Error  error
	db     *sql.DB
	debug  bool
	logger *log.Logger
}

type Param map[string]interface{}
type RawValue string

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
