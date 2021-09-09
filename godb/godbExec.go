package godb

func (db *DB) Exec(query string, args ...Param) (int64, error) {
	if len(args) == 0 {
		// 파라메터가 없는 단순쿼리
		return db.execSimple(query)
	} else {
		param := mergeMaps(args...)
		return db.execParam(query, param)
	}
}

func (db *DB) execSimple(query string) (RowsAffected int64, err error) {
	RowsAffected = -1
	if db.debug {
		db.logger.Println("QUERY : ", query)
	}

	res, err := db.db.Exec(query)
	if err == nil {
		if ra, e := res.RowsAffected(); e == nil {
			RowsAffected = ra
		}
	}
	return RowsAffected, err
}

func (db *DB) execParam(query string, args Param) (int64, error) {
	newQuery := db.replaceParamQuery(query, args)

	return db.execSimple(newQuery)
}



