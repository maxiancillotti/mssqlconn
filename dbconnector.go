package mssqlconn

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	//Database SQL Server Driver Initialization
	_ "github.com/denisenkom/go-mssqldb"
)

type DBConnector interface {
	OpenConn() *sql.DB
}

type dbConn struct {
	dbconfig *dbConnBuilder
	connOnce sync.Once
	conn     *sql.DB
}

// OpenConn opens and returns a singleton connection to the database
// previously indicated to builder.
func (db *dbConn) OpenConn() *sql.DB {

	db.connOnce.Do(func() {
		connString := db.getConnString()
		conn, err := sql.Open("mssql", connString)
		if err != nil {
			log.Panicln("Open connection to database failed: ", err.Error())
		}
		err = conn.Ping()
		if err != nil {
			log.Panicln("Connection to database failed: ", err.Error())
		}
		db.conn = conn
	})
	return db.conn
}

func (db *dbConn) getConnString() string {

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;port=%d", db.dbconfig.server, db.dbconfig.user, db.dbconfig.password, db.dbconfig.dbname, db.dbconfig.port)

	if db.dbconfig.debug {
		fmt.Printf(" server:%s\n", db.dbconfig.server)
		fmt.Printf(" port:%d\n", db.dbconfig.port)
		fmt.Printf(" user:%s\n", db.dbconfig.user)
		fmt.Printf(" password:%s\n", db.dbconfig.password)
		fmt.Printf(" dbname:%s\n", db.dbconfig.dbname)
		fmt.Printf(" connString:%s\n", connString)
	}

	return connString
}
