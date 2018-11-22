package commonsys

import (
	"os"
	"strings"
)

var env Environment
var ignoreCase = os.Getenv("COMMONSYS_IGNORE_ENV_CASE") == "1"

func init() {
	env.properties = map[string]string{}
	var record string
	for _, record = range os.Environ() {
		var parts = strings.Split(record, "=")
		var key = parts[0]
		if ignoreCase {
			key = strings.ToLower(key)
		}
		env.properties[key] = strings.Join(parts[1:], "=")
	}
}

// Environment envelops and restricts access to the enviroment variables
type Environment struct {
	properties map[string]string
}

// GetEnvironment returns the populates Environment object
func GetEnvironment() *Environment {
	return &env
}

// Get returns the value of a environment variable and if the key exists or not
func (e *Environment) Get(key string) (string, bool) {
	if ignoreCase {
		key = strings.ToLower(key)
	}
	var value string
	var ok bool
	value, ok = e.properties[key]
	return value, ok
}

// GetOrDefault returns the value of a environment variable or a default value in case the key does not exist.
func (e *Environment) GetOrDefault(key, defaultValue string) string {
	if ignoreCase {
		key = strings.ToLower(key)
	}
	var value string
	var ok bool
	value, ok = e.properties[key]
	if !ok {
		value = defaultValue
	}
	return value
}

// Set overrides the value of a key in the Environment object
func (e *Environment) Set(key, value string) {
	if ignoreCase {
		key = strings.ToLower(key)
	}
	e.properties[key] = value
}

// Remove deletes the key from the Environment object
func (e *Environment) Remove(key string) (string, bool) {
	if ignoreCase {
		key = strings.ToLower(key)
	}
	var value string
	var ok bool
	value, ok = e.properties[key]
	delete(e.properties, key)
	return value, ok
}

// Each executes a function for each key/value in the environment variables map
func (e *Environment) Each(handle func(key, value string)) {
	var key, value string
	for key, value = range e.properties {
		handle(key, value)
	}
}
