package godb

import (
	"fmt"
	"testing"

	//"time"

	//"database/sql"
	_ "github.com/lib/pq"
)
var db *DB

func init() {
    dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", "localhost", "admin", "admin", "godb")

    db, _ = Open("postgres", dbinfo)
}

func TestOpenDB(t *testing.T)  {
    dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", "localhost", "admin", "admin", "godb")

    _, err := Open("postgres", dbinfo)
    if err != nil {
        t.Errorf("Fail to connect DB : %v", err)
    }
}
