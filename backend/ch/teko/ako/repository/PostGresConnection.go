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
	dbname   = "ako"
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

func (storage *PostgresqlStorageImpl) CreateAdA(pName string, pRank string, pCompanyName string) (int, error) {
	company, err := storage.GetCompanyIdByName(pCompanyName)
	companyId := company.Id
	if err != nil {
		companyId = -1
	}
	result, err := storage.db.Exec("INSERT INTO tbl_ada (name, rank, companyId) VALUES ($1,$2,$3)", pName, pRank, companyId)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

func (storage *PostgresqlStorageImpl) CreateCompany(pName string) (int, error) {
	result, err := storage.db.Exec("INSERT INTO public.tbl_company (name) VALUES ($1)", pName)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

func (storage *PostgresqlStorageImpl) GetCompanyIdByName(pName string) (*model.Company, error) {
	var company model.Company
	row := storage.db.QueryRow("SELECT id FROM public.tbl_company WHERE name = $1", pName)
	err := row.Scan(&company.Id)
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (storage *PostgresqlStorageImpl) GetAllCompanies() []model.Company {
	var companies = make([]model.Company, 0)
	rows, err := storage.db.Query("SELECT id, name FROM public.tbl_company")
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
