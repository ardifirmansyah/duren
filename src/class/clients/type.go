package clients

type (
	Client struct {
		ID     string `db:"id"`
		Secret string `db:"secret"`
	}
)
