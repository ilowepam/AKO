package repository

import (
	"AKO/ch/teko/ako/model"
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

func (storage *PostgresqlStorageImpl) CreateAdA(pName string, pRank model.Rank) (int, error) {
	result, err := storage.db.Exec("INSERT INTO tbl_AdA (name, rank, age) VALUES ($1, $2)", pName, pRank)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

func (storage *PostgresqlStorageImpl) GetAllAdAByCompanyId(pCompanyId int) []model.AdA {
	var adas = make([]model.AdA, 0)
	rows, err := storage.db.Query("SELECT id, name, rank FROM tbl_Ada a WHERE  ")
	if err != nil {
		return adas
	}

	for rows.Next() {
		var ada model.AdA
		err = rows.Scan(&ada.Id, &ada.Name, &ada.Rank)
		if err != nil {
			log.Println("Could not read values from row")
			return adas
		}
		adas = append(adas, ada)
	}

	return adas
}

func (storage *PostgresqlStorageImpl) DeleteAdA(id int) bool {
	_, err := storage.db.Exec("DELETE FROM tbl_AdA WHERE id = $1", id)
	return err == nil
}
func (storage *PostgresqlStorageImpl) CreateCompany(pName string, pCommander model.AdA) (int, error) {
	result, err := storage.db.Exec("INSERT INTO tbl_AdA (name, rank) VALUES ($1, $2)", pName, pCommander)
	if err != nil {
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}
	return int(id), nil
}

func (storage *PostgresqlStorageImpl) GetLessonById(pId int) (*model.Lesson, error) {
	var lesson model.Lesson
	row := storage.db.QueryRow("SELECT id, task, done FROM tbl_lesson WHERE id = $1", pId)
	err := row.Scan(&lesson.Id, &lesson.Duration, &lesson.Desc)
	if err != nil {
		return nil, err
	}
	return &lesson, nil
}
