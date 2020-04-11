package clients

type (
	Repository interface {
		GetAllClients() ([]Client, error)
	}
)
