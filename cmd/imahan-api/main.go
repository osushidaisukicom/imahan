package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	c "github.com/osushidaisukicom/imahan-api/internal/config"
	"github.com/osushidaisukicom/imahan-api/internal/database"
)

var config *c.Config

func HandlePostTask(dbDriver, dsn string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := &database.TaskData{
			// TODO: request body から生成できるようにする
			DisplayName: "foo",
		}

		db, err := database.SetupDB(dbDriver, dsn)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		if _, err := database.InsertTaskData(db, data); err != nil {
			slog.Error(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			slog.Error(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func HandleTaskList(dbDriver, dsn string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var taskList []database.Task

		db, err := database.SetupDB(dbDriver, dsn)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		rows, err := database.ShowTaskData(db)
		if err != nil {
			slog.Error(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		for rows.Next() {
			var t database.Task

			// TODO: 項目が増えると記述が大変になるので、要改善
			if err := rows.Scan(&t.TaskID, &t.DisplayName, &t.CreatedAt, &t.UpdatedAt); err != nil {
				slog.Error(err.Error())
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			taskList = append(taskList, t)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(taskList); err != nil {
			slog.Error(err.Error())
		}
	}
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	_c, err := c.New()
	if err != nil {
		slog.Error("Initialize config is failed")
		return
	}

	config = _c

	dbDriver := "postgres"
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DB.Host,
		config.DB.Port,
		config.DB.User,
		config.DB.Password,
		config.DB.Name,
	)

	r := chi.NewRouter()
	r.Post("/task", HandlePostTask(dbDriver, dsn))
	r.Get("/task", HandleTaskList(dbDriver, dsn))

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", config.ServerPort),
		Handler: r,
	}

	server.ListenAndServe()
}
