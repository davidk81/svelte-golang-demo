package patientdb

import (
	"database/sql"
)

var sharedb *sql.DB

// Init initialize database connection
func Init() {
	db, err := sql.Open("postgres", "host=localhost dbname=patientdb user=docker password=docker sslmode=disable")
	if err != nil {
		panic(err)
	}
	sharedb = db

	// TODO: look into setting exec db context globally instead
	// boil.SetDB(db)
	// boil.DebugMode = true
}

// DB returns instance for patient db
func DB() *sql.DB {
	return sharedb
}

func Close() {
	sharedb.Close()
}
