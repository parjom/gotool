package godb

import (
	"fmt"
	"testing"

	//"time"

	//"database/sql"
	_ "github.com/lib/pq"
)

func TestGoDBOpen(t *testing.T)  {

    dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", "localhost", "admin", "admin", "godb")

    db, err := Open("postgres", dbinfo)
    if err != nil {
        t.Errorf("Fail to connect DB : %v", err)
    } else {

        _, _, err := db.Exec("DROP TABLE IF EXISTS test; create table test ( id int);")
        if err != nil {
            t.Errorf("Fail to Query : %v", err)
        }

        {
            query := `
            delete from test;
            `
            ra, li, err := db.Exec(query)
            fmt.Printf("RowsAffected : %v, LastInsertId : %v \n",ra, li )
            if err != nil {
                t.Errorf("Fail to Query : %v", err)
            }

        }

        {
            query := `
            DROP TABLE test; create table test ( id int);
            insert into test(id) values ( 101 );
            insert into test(id) values ( 102 );
            insert into test(id) values ( 103 );
            insert into test(id) values ( 104 );
            insert into test(id) values ( 105 );
            insert into test(id) values ( 106 );
            insert into test(id) values ( 107 );
            insert into test(id) values ( 108 );
            insert into test(id) values ( 109 );
            insert into test(id) values ( 110 );
            insert into test(id) values ( 111 );
            `
            ra, li, err := db.Exec(query)
            fmt.Printf("RowsAffected : %v, LastInsertId : %v \n",ra, li )
            if err != nil {
                t.Errorf("Fail to Query : %v", err)
            }
        }

        // for i:=0; i<100; i++ {
        //     _, _, err := db.Exec("insert into test(id) values ( $1 )", i)
        //     if err != nil {
        //         t.Errorf("Fail to Query : %v", err)
        //     }
        //     time.Sleep(10*time.Second)
        // }
        //time.Sleep(10*time.Second)
    }

}