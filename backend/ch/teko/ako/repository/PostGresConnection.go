package repository

import (
	"AKO/ch/teko/ako/model"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
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

func (storage *PostgresqlStorageImpl) GetAllCompanies() []model.Company {
	var companies = make([]model.Company, 0)
	rows, err := storage.db.Query("SELECT id, name FROM tbl_company")
	if err != nil {
		return companies
	}

	for rows.Next() {
		var company model.Company
		err = rows.Scan(&company.Id, &company.Name)
		if err != nil {
			log.Println("Could not read values from row")
			return companies
		}
		companies = append(companies, company)
	}
	return companies
}
