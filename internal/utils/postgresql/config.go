// provides a standard way to get a postgresql database connection string from the env

package postgresql

import (
	"fmt"

	"github.com/GOOFR-Group/store-back-end/internal/utils/env"
)

// Defaults supplies your environment defaults via k/v pairs
type Defaults = env.Defaults

// Configuration stores our configuration values
type Configuration struct {
	env.DbConnectionConfiguration
	Prefix string
}

// String is a log friendly string version of Configuration (no passwords)
func (config Configuration) String() string {
	if config.Prefix == "" {
		return fmt.Sprintf("%v", config.DbConnectionConfiguration)
	}
	return fmt.Sprintf("(%s) %v", config.Prefix, config.DbConnectionConfiguration)
}

// LoadConfigurationFromEnv load configuration from env
func LoadConfigurationFromEnv(defaults Defaults) (dbconfig Configuration) {
	dbconfig.DbConnectionConfiguration = env.GetDbConfig(defaults)
	return
}

// LoadConfigurationFromEnvWithPrefix load configuration from env for a specific subsystem
func LoadConfigurationFromEnvWithPrefix(prefix string, defaults Defaults) (dbconfig Configuration) {
	dbconfig.Prefix = prefix
	dbconfig.DbConnectionConfiguration = env.GetDbConfigFor(prefix, defaults)
	return
}
