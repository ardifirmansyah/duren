package class

import (
	"github.com/ardifirmansyah/duren/src/class/clients"
	"github.com/ardifirmansyah/duren/src/common/database"
)

func Init(db *database.DBConnection) {
	clients.New(db).Init()
}
