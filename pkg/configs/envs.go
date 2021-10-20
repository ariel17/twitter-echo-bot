package configs

import "os"

const productionEnv = "production"
const environmentKey = "ENVIRONMENT"

var environment string

// IsProduction checks the environment name and indicates if it is production or
// not.
func IsProduction() bool {
	return environment == productionEnv
}

func loadEnvs() {
	environment = os.Getenv(environmentKey)
}

func init() {
	loadEnvs()
}