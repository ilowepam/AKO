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
	company, err := storage.GetCompanyByName(pCompanyName)
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

func (storage *PostgresqlStorageImpl) GetLessonByName(pName string) *model.Lesson {
	var lesson model.Lesson
	row := storage.db.QueryRow("SELECT id, name FROM public.tbl_lesson WHERE name = $1", pName)
	err := row.Scan(&lesson.Id, &lesson.Name)
	if err != nil {
		return nil
	}
	return &lesson
}

func (storage *PostgresqlStorageImpl) CreateLessonStatus(pAdaId int, pLessonId int, pStatus bool) (bool, error) {
	result, err := storage.db.Exec("INSERT INTO public.tbl_lesson_status (adaid, lessonid, status) VALUES ($1,$2,$3)", pAdaId, pLessonId, pStatus)
	if err != nil {
		return false, err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return false, err
	}
	return true, nil
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

func (storage *PostgresqlStorageImpl) GetLessonStatus(pAdaId int, pLessonId int) (bool, error) {
	var lessStatus bool
	row := storage.db.QueryRow("SELECT status FROM public.tbl_lesson_status WHERE adaId = $1 AND lessonId = $2", pAdaId, pLessonId)
	err := row.Scan(&lessStatus)
	if err != nil {
		return lessStatus, err
	}
	return lessStatus, nil
}

func (storage *PostgresqlStorageImpl) GetCompanyById(pId int) (*model.Company, error) {
	var company model.Company
	row := storage.db.QueryRow("SELECT id, name FROM public.tbl_company WHERE id = $1", pId)
	err := row.Scan(&company.Id, &company.Name)
	if err != nil {
		return nil, err
	}
	return &company, err
}

func (storage *PostgresqlStorageImpl) GetCompanyByName(pName string) (*model.Company, error) {
	var company model.Company
	row := storage.db.QueryRow("SELECT id, name FROM public.tbl_company WHERE name = $1", pName)
	err := row.Scan(&company.Id, &company.Name)
	if err != nil {
		return nil, err
	}
	return &company, nil
}

func (storage *PostgresqlStorageImpl) GetAllAdAs() []model.AdA {
	var adas = make([]model.AdA, 0)
	rows, err := storage.db.Query("SELECT id, name, rank, companyId FROM public.tbl_ada")
	if err != nil {
		return adas
	}
	for rows.Next() {
		var ada model.AdA
		err = rows.Scan(&ada.Id, &ada.Name, &ada.Rank, &ada.CompanyId)
		if err != nil {
			log.Println("Could not read values from row")
			return adas
		}
		adas = append(adas, ada)
	}
	return adas
}

func (storage *PostgresqlStorageImpl) GetAllLessons() []model.Lesson {
	var lessons = make([]model.Lesson, 0)
	rows, err := storage.db.Query("SELECT id, name FROM public.tbl_lesson")
	if err != nil {
		return lessons
	}
	for rows.Next() {
		var lesson model.Lesson
		err = rows.Scan(&lesson.Id, &lesson.Name)
		if err != nil {
			log.Println("Could not read values from row")
			return lessons
		}
		lessons = append(lessons, lesson)
	}
	return lessons
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

func (storage *PostgresqlStorageImpl) CreateLesson(pName string) (int, error) {
	result, err := storage.db.Exec("INSERT INTO public.tbl_lesson (name) VALUES ($1)", pName)
	if err != nil {
		log.Println("Could not create lesson")
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Println("Could not read last inserted id")
		return -1, err
	}
	return int(id), nil
}
