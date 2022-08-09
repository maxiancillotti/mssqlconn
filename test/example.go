// This is only a handy package to test an actual connection to a database.
// See the comments and change the data below.
package main

import (
	"log"

	"github.com/maxiancillotti/mssqlconn"
)

var conn = mssqlconn.NewBuilder().
	// COMPLETE WITH VALID INPUTS
	SetHostname("localhost").
	SetPort(1433).
	SetInstance("SQLEXPRESS").
	SetDatabaseName("DBname").
	SetCredentials("sa", "password").
	EnableDebug().
	Build().
	OpenConn()

func main() {
	defer conn.Close()

	rows, err := conn.Query("select * from TABLE") // COMPLETE WITH A VALID QUERY

	if err != nil {
		log.Fatal("SP Exec failed:", err.Error())
	}

	defer rows.Close()
	for rows.Next() {

		// CHANGE THE VARIABLES TO SCAN ACTUAL DATABASE OUTPUT
		var id int
		var nationalityID int
		var nationalIDType int
		var nationalID string
		var name string
		var surname string

		err = rows.Scan(&id, &nationalityID, &nationalIDType, &nationalID, &name, &surname) // ALSO HERE
		if err != nil {
			log.Fatal("SP result rows scan failed:", err.Error())
		}

		log.Println(id, nationalityID, nationalIDType, nationalID, name, surname) // AND HERE
	}
	err = rows.Err()
	if err != nil {
		log.Fatal("Rows scan returned an error:", err.Error())
	}
}
