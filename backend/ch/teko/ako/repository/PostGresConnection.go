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

func (storage *PostgresqlStorageImpl) GetLessonById(id int) (*model.Lesson, error) {
	var lesson model.Lesson
	row := storage.db.QueryRow("SELECT id, task, done FROM tbl_lesson WHERE id = $1", id)
	err := row.Scan(&lesson.Id, &lesson.Time, &lesson.Desc)
	if err != nil {
		return nil, err
	}
	return &lesson, nil
}
