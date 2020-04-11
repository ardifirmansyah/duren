package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ardifirmansyah/duren/src/class"
	"github.com/ardifirmansyah/duren/src/class/clients"
	"github.com/ardifirmansyah/duren/src/common/database"
)

var (
	db *database.DBConnection
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Worlds! %s, %s", time.Now(), string(getClients()))
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
