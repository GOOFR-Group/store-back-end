// supplies a standardized way to get a database connection string from env vars

package env

import (
	"strings"

	"github.com/goofr-group/store-back-end/internal/utils/format"
)

const (
	password = "password"
	host     = "host"
	dbname   = "dbname"
	user     = "user"
	port     = "port"
	sslmode  = "sslmode"
	redacted = "[REDACTED]"
)

// DbConnectionConfiguration stores our configuration values for a connection string
type DbConnectionConfiguration struct {
	DbHost    string
	DbOpts    string
	DbName    string
	DbUser    string
	DbPass    string
	DbPort    string
	DbSSLMode string

	DbConnectionString string
}

// load configuration from env for the root (DB_XXX)
func GetDbConfig(defaults Defaults) (config DbConnectionConfiguration) {
	return GetDbConfigFor("", defaults)
}

// load configuration using a prefix for all settings
func GetDbConfigFor(prefix string, defaults Defaults) (config DbConnectionConfiguration) {

	// allow a full override
	config.DbConnectionString = GetEnvOrDefaultMap(GetSubVar(prefix, "DB_CONNECTION"), defaults)
	if config.DbConnectionString != "" {
		return
	}

	// load from env or defaults or global defaults
	config.DbOpts = GetEnvOrDefaultMap(GetSubVar(prefix, "DB_OPTIONS"), defaults)
	config.DbHost = GetEnvOrDefaultCascade(GetSubVar(prefix, "DB_HOST"), defaults, "/var/run/postgresql")
	config.DbName = GetEnvOrDefaultMap(GetSubVar(prefix, "DB_NAME"), defaults)
	config.DbUser = GetEnvOrDefaultMap(GetSubVar(prefix, "DB_USER"), defaults)
	config.DbPass = GetEnvOrDefaultMap(GetSubVar(prefix, "DB_PASSWORD"), defaults)
	config.DbPort = GetEnvOrDefaultMap(GetSubVar(prefix, "DB_PORT"), defaults)
	config.DbSSLMode = GetEnvOrDefaultMap(GetSubVar(prefix, "DB_SSLMODE"), defaults)

	// construct connection string
	config.DbConnectionString = format.Assemble(config.DbConnectionString, config.DbOpts)
	config.DbConnectionString = format.AssembleFromKeyValue(config.DbConnectionString, host, config.DbHost)
	config.DbConnectionString = format.AssembleFromKeyValue(config.DbConnectionString, user, config.DbUser)
	config.DbConnectionString = format.AssembleFromKeyValue(config.DbConnectionString, password, config.DbPass)
	config.DbConnectionString = format.AssembleFromKeyValue(config.DbConnectionString, dbname, config.DbName)
	config.DbConnectionString = format.AssembleFromKeyValue(config.DbConnectionString, port, config.DbPort)
	config.DbConnectionString = format.AssembleFromKeyValue(config.DbConnectionString, sslmode, config.DbSSLMode)

	return
}

// this returns a log-safe string version of itself (no passwords)
func (config DbConnectionConfiguration) String() string {

	// find start of password
	i1 := strings.Index(config.DbConnectionString, password)
	if i1 == -1 {
		// no password, nothing to cut out
		return config.DbConnectionString
	}

	// move to start of the actual password itself (after = char)
	i1 += len(password) + 1
	i2 := strings.IndexByte(config.DbConnectionString[i1:], ' ')
	if i2 == -1 {
		// last element in the string - so cut to end of string
		i2 = len(config.DbConnectionString)
	} else {
		// adjust for our offset to find absolute start of password
		i2 += i1
	}

	// replace the password with [REDACTED]
	return config.DbConnectionString[:i1] + redacted + config.DbConnectionString[i2:]
}
