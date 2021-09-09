package godb

import (
	"testing"

	//"time"

	"database/sql"
	_ "github.com/lib/pq"
)

func TestQuery(t *testing.T)  {
    if db == nil {
        t.Errorf("DB Connection Error")
    } else {
        var (
            err error
            rowsAffected int64
            rows *sql.Rows
            query string
        )

        db.SetDebug(true)

        // 각각의 컬럼을 테스트할 임시 테이블을 생성하고,
        query = `
        DROP TABLE IF EXISTS public.test_data_type;
        CREATE TABLE IF not exists public.test_data_type (
            data_uid uuid NOT NULL,
            data_boolean boolean,
            data_char char(10),
            data_varchar varchar(100),
            data_text text,
            data_bytea bytea,
            data_geom geometry(geometry, 4326),
            data_state state,
            data_timestamptz timestamptz,
            data_timetz timetz,
            data_timestamp timestamp,
            data_date date,
            data_time time,
            data_smallint smallint,
            data_integer integer,
            data_bigint bigint,
            data_real real,
            data_float float,
            data_bit BIT(10),
            data_xml xml,
            data_json json,
            data_jsonb jsonb,
            data_integer_array integer[],
            data_varchar_array varchar[],
            data_varchar_array2 varchar[][],
            data_varchar_null varchar,
            CONSTRAINT test_data_type_pk PRIMARY KEY (data_uid)
        );
        CREATE index IF not exists sidx_test_data_type_geom ON public.test_data_type USING gist (data_geom);
        insert into public.test_data_type (
         data_uid
        ,data_boolean
        ,data_char
        ,data_varchar
        ,data_text
        ,data_bytea
        ,data_geom
        ,data_state
        ,data_timestamptz
        ,data_timetz
        ,data_timestamp
        ,data_date
        ,data_time
        ,data_smallint
        ,data_integer
        ,data_bigint
        ,data_real
        ,data_float
        ,data_bit
        ,data_xml
        ,data_json
        ,data_jsonb
        ,data_integer_array
        ,data_varchar_array
        ,data_varchar_array2
        ,data_varchar_null
        )
        values (
         uuid_generate_v4()
        ,true
        ,'qwer'
        ,'abcdefg'
        ,'hijklmn'
        ,decode('f4114e1dc44b78a5a36b6f3f482d6fc5c24e83e9eef9','hex')
        ,ST_GeomFromText('POINT(126.8901849 37.4861311)', 4326)
        ,'ok'
        ,now()
        ,now()
        ,now()
        ,now()
        ,now()
        ,2
        ,4
        ,8
        ,0.4
        ,0.8
        ,'0111110000'
        ,'<xml></xml>'
        ,'{"bar": "baz",      "balance": 7.77, "active":false}'::json
        ,'{"bar": "baz", "balance": 7.77, "active":false}'::jsonb
        ,'{10000, 10001, 10002, 10003}'
        ,'{"meeting", "lunch"}'
        ,'{{"meeting1", "lunch1"}, {"meeting2", "lunch2"}}'
        ,null
        );

        `
        rowsAffected, err = db.Exec(query)
        if err != nil {
            t.Errorf("Fail to Query : %v", err)
        } else {
            t.Logf("RowsAffected : %v\n", rowsAffected)
        }

        // 여기서 부터 실질적인 테스트
        query = `select * from public.test_data_type`
        rows, err = db.Query(query)
        if err != nil {
            t.Errorf("Fail to Query : %v", err)
        } else {
            t.Logf("RowsAffected : %v\n", rowsAffected)
        }

        _, err = ScanRows(rows)

    }

}
