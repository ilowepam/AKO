package repository

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "0.0.0.0"
	port     = 5432
	user     = "postgres"
	password = "secret"
	dbname   = "reminder"
)

type PostgresqlStorageImpl struct {
	db *sql.DB
}

func NewPostgresStorageImpl() *PostgresqlStorageImpl {
	postgresqlConnection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", postgresqlConnection)
	if err != nil {
		log.Fatalln(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}

	return &PostgresqlStorageImpl{db: db}
}
