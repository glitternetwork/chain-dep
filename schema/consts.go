package schema

import "strings"

const (
	DefaultStoreDatabaseName = "glitter"
)

var (
	ReservedDatabases = []string{
		"metrics_schema",
		"information_schema",
		"performance_schema",
		"mysql",
		"test",
	}
)

func IsReservedDatabase(name string) bool {
	for _, v := range ReservedDatabases {
		if v == strings.ToLower(name) {
			return true
		}
	}
	return false
}
