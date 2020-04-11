package clients

import (
	"github.com/jmoiron/sqlx"

	"github.com/ardifirmansyah/duren/src/common/database"
)

const (
	getAllClientsQuery = `SELECT id, secret FROM clients`
)

var (
	localRepo *PSQLRepository
)

type (
	PSQLRepository struct {
		masterDB   database.DBService
		statements preparedStatements
	}

	preparedStatements struct {
		GetAllClients *sqlx.Stmt
	}
)

func New(db *database.DBConnection) *PSQLRepository {
	if localRepo != nil {
		return localRepo
	}
	localRepo = &PSQLRepository{
		masterDB: db.Master,
	}

	return localRepo
}

func (repo *PSQLRepository) Init() {
	repo.statements = preparedStatements{
		GetAllClients: repo.masterDB.Preparex(getAllClientsQuery),
	}
}

func (repo *PSQLRepository) GetAllClients() ([]Client, error) {
	cs := []Client{}

	err := repo.statements.GetAllClients.Select(&cs)
	if err != nil {
		return nil, err
	}

	return cs, nil
}
