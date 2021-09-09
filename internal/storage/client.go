package storage

// Client model
type Client struct {
	tableName struct{} `pg:"clients"`

	ID         int    `pg:"id"`          // ID
	ClientName string `pg:"client_name"` // Наименование
	UrAdr      string `pg:"ur_adr"`      // Юридический адрес
}
