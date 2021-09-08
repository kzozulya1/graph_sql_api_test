package config

// Configuration is main app config
type Configuration struct {
	ListenAddr string `envconfig:"LISTEN_ADDR"      default:"0.0.0.0:2121"`
	DBConn     string `envconfig:"DB_CONN"          default:"postgres://postgres:postgres@127.0.0.1:5432/mesh_group?sslmode=disable"`

	DBSQLQueryLog bool `envconfig:"DB_SQL_QUERY_LOG" default:"false"`
}
