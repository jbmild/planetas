package Utils

import (
    "database/sql"
   	_ "github.com/go-sql-driver/mysql"
   	"fmt"
)

func GetConnection(user string, password string, host string, port string, database string) *sql.DB{

	t := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	db, err := sql.Open("mysql", t)
	if err != nil {
        panic(err.Error())
    }

    return db
}

func ExecuteQuery(db *sql.DB, query string) bool{
	q, err := db.Query(query)
    if err != nil {
        return false
    }
    q.Close()
    return true
}