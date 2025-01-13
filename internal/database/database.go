package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/osushidaisukicom/imahan-api/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Task struct {
	*models.Task
	db *sql.DB
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

func InsertTaskData(ctx context.Context, db *sql.DB, data *TaskData) (*models.Task, error) {
	task := &models.Task{
		DisplayName: data.DisplayName,
	}

	err := task.Insert(ctx, db, boil.Infer())
	if err != nil {
		return nil, err
	}

	return task, nil
}

func ShowTaskData(ctx context.Context, db *sql.DB) (models.TaskSlice, error) {
	result, err := models.Tasks().All(ctx, db)
	if err != nil {
		return nil, err
	}

	return result, nil
}
