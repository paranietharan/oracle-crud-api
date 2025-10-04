package db

import (
	"database/sql"
	"fmt"
	"log"
	"oracle-crud-api/internal/config"

	_ "github.com/godror/godror"
)

func ConnectToOracleDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(`user="%s" password="%s" connectString="%s"`,
		cfg.Username, cfg.Password, cfg.ConnectionString,
	)
	db, err := sql.Open("godror", dsn)
	if err != nil {
		log.Fatal("Error connecting to Oracle DB: ", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging Oracle DB: ", err)
		return nil, err
	}

	fmt.Println("Connected to Oracle Database!")
	return db, nil
}
