package database

import "database/sql"

func InitDB() *sql.DB {
	connectionString := "root:$ystem123@tcp(localhost:3306)/nortwind"

	databaseConnection, error := sql.Open("mysql", connectionString)
	if error != nil {
		panic(error.Error())
	}

	return databaseConnection
}
