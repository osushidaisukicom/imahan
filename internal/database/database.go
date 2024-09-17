package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	TaskID      uuid.UUID
	DisplayName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TaskData struct {
	DisplayName string
}

func SetupDB(dbDriver string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		return nil, err
	}
	return db, err
}

func GenUUID() uuid.UUID {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return uuid
}

func InsertTaskData(db *sql.DB, data *TaskData) (sql.Result, error) {
	result, err := db.Exec(`INSERT INTO task_list (display_name) VALUES ($1)`, data.DisplayName)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func ShowTaskData(db *sql.DB) (*sql.Rows, error) {
	result, err := db.Query(`SELECT * FROM task_list`)
	if err != nil {
		return result, err
	}

	return result, nil
}
