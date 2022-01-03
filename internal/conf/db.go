package conf

import "github.com/GOOFR-Group/store-back-end/internal/utils/postgresql"

const (
	DBHost    = "DB_HOST"
	DBName    = "DB_NAME"
	DBUser    = "DB_USER"
	DBPort    = "DB_PORT"
	DBSSLMode = "DB_SSLMODE"
)

const (
	DefaultDBHost  = "localhost"
	DefaultDBName  = "store_db"
	DefaultDBUser  = "goofr"
	DefaultDBPort  = "5432"
	DefaultSSLMode = "enable"
)

var dbConfiguration postgresql.Configuration

// InitDB starts the database configuration
func InitDB() {
	dbConfiguration = postgresql.LoadConfigurationFromEnv(postgresql.Defaults{
		DBHost:    DefaultDBHost,
		DBName:    DefaultDBName,
		DBUser:    DefaultDBUser,
		DBPort:    DefaultDBPort,
		DBSSLMode: DefaultSSLMode,
	})
}

// GetDbConnectionString retrieves the configured database connection string
func GetDbConnectionString() string {
	return dbConfiguration.DbConnectionString
}
