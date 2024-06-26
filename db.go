package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type Database interface {
	GetCars(page int, pageSize int, make, model string, year int) ([]*Car, error)
	DeleteCarByID(id int) error
	UpdateCarByID(id int, car *Car) error
	AddCars(regNums []string) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	host, _ := os.LookupEnv("HOST")
	portStr, _ := os.LookupEnv("DB_PORT")
	port, _ := strconv.Atoi(portStr)
	user, _ := os.LookupEnv("USERNAME")
	password, _ := os.LookupEnv("PASSWORD")
	dbname, _ := os.LookupEnv("DB_NAME")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	store := &PostgresStore{db: db}
	err = initTables(store)
	if err != nil {
		return nil, err
	}
	return store, nil
}

func initTables(s *PostgresStore) error {
	peopleQuery := `CREATE TABLE IF NOT EXISTS people (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    surname VARCHAR(255) NOT NULL,
    patronymic VARCHAR(255)
);`
	_, err := s.db.Query(peopleQuery)
	if err != nil {
		return err
	}
	carQuery := `CREATE TABLE IF NOT EXISTS cars (
    id SERIAL PRIMARY KEY,
    regNum VARCHAR(20) NOT NULL,
    mark VARCHAR(255) NOT NULL,
    model VARCHAR(255) NOT NULL,
    year INTEGER,
    owner_id INTEGER REFERENCES people(id)
);
  `
	_, err = s.db.Query(carQuery)
	return err
}

func (s *PostgresStore) GetCars(page int, pageSize int, make string, model string, year int) ([]*Car, error) {
	var cars []*Car
	offset := (page - 1) * pageSize

	query := "SELECT * FROM cars"

	conditions := []string{}
	if make != "" {
		conditions = append(conditions, fmt.Sprintf("make = '%s'", make))
	}
	if model != "" {
		conditions = append(conditions, fmt.Sprintf("model = '%s'", model))
	}
	if year != 0 {
		conditions = append(conditions, fmt.Sprintf("year = %d", year))
	}
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	query += fmt.Sprintf(" LIMIT %d OFFSET %d", pageSize, offset)

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		car, err := ScanIntoCar(rows)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	return cars, nil
}

func (s *PostgresStore) DeleteCarByID(id int) error {
	_, err := s.db.Query(`DELETE FROM cars WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) UpdateCarByID(id int, car *Car) error {
	query := `
        UPDATE cars
        SET reg_num = $1, mark = $2, model = $3, year = $4, owner_name = $5, owner_surname = $6, owner_patronymic = $7
        WHERE id = $8
    `
	_, err := s.db.Exec(
		query,
		car.RegNum,
		car.Mark,
		car.Model,
		car.Year,
		car.Owner.Name,
		car.Owner.Surname,
		car.Owner.Patronymic,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) AddCars(regNums []string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	stmt, err := tx.Prepare("INSERT INTO cars (reg_num) VALUES ($1)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, regNum := range regNums {
		_, err := stmt.Exec(regNum)
		if err != nil {
			return err
		}
	}

	return nil
}

func ScanIntoCar(rows *sql.Rows) (*Car, error) {
	car := new(Car)
	err := rows.Scan(
		&car.Owner,
		&car.Mark,
		&car.Year,
		&car.RegNum,
	)
	return car, err
}
