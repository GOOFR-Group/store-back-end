package env

import (
	"fmt"
	"os"

	"github.com/goofr-group/store-back-end/internal/utils/format"
)

// Defaults supplies your environment defaults via k/v pairs
type Defaults = map[string]string

// GetSubVar returns a variable name in the form of KEY_SUBKEY
func GetSubVar(key, subkey string) string {
	return format.AssembleWith(key, subkey, "_")
}

// GetEnvOrPanic returns the specified environment variable's contents, or panics
func GetEnvOrPanic(key string) (value string) {
	value, found := os.LookupEnv(key)
	if !found {
		panic(fmt.Sprintf("Required environment variable '%s' not set", key))
	}
	return
}

// GetEnvOrDefault returns the specified environment variable's contents, or the specified default value if that env var is not present
func GetEnvOrDefault(key, defvalue string) (value string) {
	value, found := os.LookupEnv(key)
	if !found {
		value = defvalue
	}
	return
}

// GetEnvOrDefaultMap returns the specified environment variable's contents, or the same value in the defaults map
func GetEnvOrDefaultMap(key string, defaults Defaults) (value string) {
	value, found := os.LookupEnv(key)
	if !found {
		value = defaults[key]
	}
	return
}

// GetEnvOrDefaultCascade returns the specified environment variable's contents, or the same value in the defaults map, or a global default if nothing in map for this key
func GetEnvOrDefaultCascade(key string, defaults Defaults, global string) (value string) {
	value, found := os.LookupEnv(key)
	if !found {
		value, found = defaults[key]
		if !found {
			value = global
		}
	}
	return
}

// CreateEnvValueFromEnvFile creates an environment variable from the contents of another variable that points to a file.
// If `ignoreIfExists` is set to true and the environment variable already exists, its creation will be ignored.
func CreateEnvValueFromEnvFile(envValue, envFile string, ignoreIfExists bool) {
	if ignoreIfExists {
		if _, ok := os.LookupEnv(envValue); ok {
			return
		}
	}

	file, ok := os.LookupEnv(envFile)
	if !ok {
		return
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return
	}

	os.Setenv(envValue, string(data))
}
