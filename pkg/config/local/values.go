package local

import (
	"os"

	"github.com/rafaelsanzio/go-core/pkg/config/key"
)

var Values = map[key.Key]string{
	key.MongoDBName:     getDefaultOrEnvVar("go-core", "MONGO_DATABASE"),
	key.MongoDBPassword: "",
	key.MongoDBUsername: "",
	key.MongoURI:        getDefaultOrEnvVar("mongodb://localhost:27017", "MONGO_URI"),
	key.Region:          "us-east-1",
}

// Some of the db fields are set via env var in the makefile, so this optionally uses those to prevent test failures in jenkins
func getDefaultOrEnvVar(dfault, envVar string) string {
	val := os.Getenv(envVar)
	if val != "" {
		return val
	}
	return dfault
}
