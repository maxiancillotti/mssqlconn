package main

import (
	"log"

	"github.com/maxiancillotti/mssqlconn"
)

var conn = mssqlconn.NewBuilder().
	SetServer("localhost\\SQLEXPRESS").
	SetPort(1433).
	SetDatabaseName("Barocco_Wallet").
	SetCredentials("sa", "36801583").
	EnableDebug().
	Build().
	OpenConn()

func main() {
	rows, err := conn.Query("select * from Customers")

	if err != nil {
		log.Fatal("SP Exec failed:", err.Error())
	}

	defer rows.Close()
	for rows.Next() {

		var id int
		var nationalityID int
		var nationalIDType int
		var nationalID string
		var name string
		var surname string

		err = rows.Scan(&id, &nationalityID, &nationalIDType, &nationalID, &name, &surname)
		if err != nil {
			log.Fatal("SP result rows scan failed:", err.Error())
		}

		log.Println(id, nationalityID, nationalIDType, nationalID, name, surname)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal("Rows scan returned an error:", err.Error())
	}
}
