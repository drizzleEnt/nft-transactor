package db

import (
	"database/sql"
	"fmt"
)

func ConnectDB(user, password, host string, port int, dbname string) (*sql.DB, error) {
	dbDNS := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgers", dbDNS)
	if err != nil {
		return nil, fmt.Errorf("Failed open db connection: %s", err.Error())
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("Failed ping db: %s", err.Error())
	}

	return db, nil
}
