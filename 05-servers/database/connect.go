package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var _db *sql.DB = nil
var _db_closer *time.Timer

func Use() *sql.DB {

	if _db == nil {
		// do the connection thing
		var err error
		_db, err = sql.Open("postgres", "postgresql://fluffycat:s3cret@172.19.0.2:5432/heavycake?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}

		// setups to close the timer
		_db_closer = time.AfterFunc(5*time.Second, func() {
			_db.Close()
			_db = nil
		})
	} else {
		// já existe uma instância ativa, só reseta o timer
		_db_closer.Reset(5 * time.Second)
	}

	return _db
}
