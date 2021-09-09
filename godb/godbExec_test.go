package godb

import (
	"testing"

	//"time"

	//"database/sql"
	_ "github.com/lib/pq"
)

func TestExec(t *testing.T)  {
    if db == nil {
        t.Errorf("DB Connection Error")
    } else {
        var (
            err error
            rowsAffected int64
            query string
        )

        db.SetDebug(true)

        query = "create table IF NOT EXISTS test ( id serial, data int);"
        rowsAffected, err = db.Exec(query)
        if err != nil {
            t.Errorf("Fail to Query : %v", err)
        } else {
            t.Logf("RowsAffected : %v\n", rowsAffected)
        }

        query = `delete from test;`
        rowsAffected, err = db.Exec(query)
        if err != nil {
            t.Errorf("Fail to Query : %v", err)
        } else {
            t.Logf("RowsAffected : %v\n", rowsAffected)
        }

        query = `
        insert into test( data ) values ( 101 );
        `
        rowsAffected, err = db.Exec(query)
        if err != nil {
            t.Errorf("Fail to Query : %v", err)
        } else {
            t.Logf("RowsAffected : %v\n", rowsAffected)
        }

        query = `
        insert into test( data ) values ( 102 );
        insert into test( data ) values ( 103 );
        insert into test( data ) values ( 104 );
        insert into test( data ) values ( 105 );
        insert into test( data ) values ( 106 );
        `
        rowsAffected, err = db.Exec(query)
        if err != nil {
            t.Errorf("Fail to Query : %v", err)
        } else {
            t.Logf("RowsAffected : %v\n", rowsAffected)
        }

        query = `
        insert into test( data ) values ( :data );
        `
        rowsAffected, err = db.Exec(query, Param{
            "data": 201,
        })
        if err != nil {
            t.Errorf("Fail to Query : %v", err)
        } else {
            t.Logf("RowsAffected : %v\n", rowsAffected)
        }

    }

}
