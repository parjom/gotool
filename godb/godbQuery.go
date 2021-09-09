package godb

import (
    "database/sql"
    "log"
)

func (db *DB) Query(query string, args ...Param) (*sql.Rows, error) {
    if len(args) == 0 {
        // 파라메터가 없는 단순쿼리
        return db.querySimple(query)
    } else {
        param := mergeMaps(args...)
        return db.queryParam(query, param)
    }
}

func (db *DB) querySimple(query string) (*sql.Rows, error) {

    if db.debug {
        db.logger.Println("QUERY : ", query)
    }

    return db.db.Query(query)
}

func (db *DB) queryParam(query string, args Param) (*sql.Rows, error) {
    newQuery := db.replaceParamQuery(query, args)

    return db.querySimple(newQuery)
}

func ScanRows(rows *sql.Rows) ([]map[string]interface{}, error) {
    result := []map[string]interface{}{}

    colTypes, err := rows.ColumnTypes()
    if err != nil {
        // 컬럼타입을 가져오는중 에러 발생
        return nil, err
    }

    for _,s := range colTypes {
        log.Println("cols type:", s.Name(), s.DatabaseTypeName(), s.ScanType());
    }

    return result, nil
}



// data_uid				UUID				interface{}
// data_boolean			BOOL				bool
// data_char			BPCHAR				interface{}
// data_varchar			VARCHAR				string
// data_text			TEXT				string
// data_bytea			BYTEA				[]uint8
// data_geom								interface{}
// data_state								interface{}
// data_timestamptz		TIMESTAMPTZ			time.Time
// data_timetz			TIMETZ				time.Time
// data_timestamp		TIMESTAMP			time.Time
// data_date			DATE				time.Time
// data_time			TIME				time.Time
// data_smallint		INT2				int16
// data_integer			INT4				int32
// data_bigint			INT8				int64
// data_real			FLOAT4				interface{}
// data_float			FLOAT8				interface{}
// data_bit				BIT					interface{}
// data_xml				XML					interface{}
// data_json			JSON				interface{}
// data_jsonb			JSONB				interface{}
// data_integer_array	_INT4				interface{}
// data_varchar_array	_VARCHAR			interface{}
// data_varchar_array2	_VARCHAR			interface{}
// data_varchar_null	VARCHAR 			string
