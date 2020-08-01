package patientdb

// line below enables 'go generate' cmd to update orm stubs
//go:generate sqlboiler --wipe psql

import (
	"database/sql"
)

var sharedb *sql.DB

// Init initialize database connection
func Init(dbConn string) {
	db, err := sql.Open("postgres", dbConn)
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
