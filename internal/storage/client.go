package storage

// ClientTablename holds persistent tablename for client
// const ClientTablename = "clients"

// Client model
type Client struct {
	tableName struct{} `pg:"clients"`

	ID          int    `pg:"id"`            // ID
	ClientSAPID string `pg:"client_sap_id"` // Код дебитора (ЮЛ/ИП)
	ClientName  string `pg:"client_name"`   // Наименование
	UrAdr       string `pg:"ur_adr"`        // Юридический адрес
	OGRN        string `pg:"ogrn"`          // ОГРН
	INN         string `pg:"inn"`           // ИНН
	KPP         string `pg:"kpp"`           // КПП
}
