package main

import (
	"encoding/json"
	"net/http"

	"github.com/ardifirmansyah/duren/src/class"
	"github.com/ardifirmansyah/duren/src/class/clients"
	"github.com/ardifirmansyah/duren/src/common/database"
)

var (
	db *database.DBConnection
)

func greet(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"database": func() string {
			if err := db.Master.Status(); err != nil {
				return err.Error()
			}
			return "database is connected"
		}(),
		"data": string(getClients()),
	}

	b, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(b)
}

func getClients() []byte {
	repo := clients.New(db)

	data, err := repo.GetAllClients()
	if err != nil {
		return []byte(err.Error())
	}

	if b, err := json.Marshal(data); err == nil {
		return b
	}
	return []byte(err.Error())
}

func main() {
	// init database connection
	db = database.GetDatabase()

	// init class
	class.Init(db)

	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
