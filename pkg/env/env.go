package env

import (
	"os"

	"github.com/joho/godotenv"
)

var Env map[string]string

func GetEnv(key, def string) string {
	if val, ok := Env[key]; ok {
		return val
	}
	return def
}

func SetupEnvFile() {
	envFile := ".env"
	var err any
	Env, err = godotenv.Read(envFile)
	if err != nil {
		defaultVariables := []string{"DB_USER", "DB_PASSWORD", "DB_PORT", "DB_NAME"}
		for _, key := range defaultVariables {
			if value := os.Getenv(key); value != "" {
				Env[key] = value
			}
		}
	}

}
