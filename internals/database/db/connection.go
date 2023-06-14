package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/cptvictor95/go-todos/internals/configs"
	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	conn, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Printf("error connecting to db: %v", err)

		return nil, err
	}

	err = conn.Ping()

	return conn, err
}
