package env

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var Env map[string]string

func GetEnv(key, def string) string {
	if val, ok := Env[key]; ok {
		return val
	}
	return def
}

func EnvLocal()  {
	envFile := ".env"
	var err any
	Env, err = godotenv.Read(envFile)
	if err != nil {
		panic(err)
	}
}

func EnvProduction()  {
	defaultVariables := []string{"DB_USER", "DB_PASSWORD", "DB_PORT", "DB_NAME"}
	Env = make(map[string]string)
	for _, key := range defaultVariables {
		if value := os.Getenv(key); value != "" {
			Env[key] = value
		}
	}
}

func SetupEnvFile() {
	appEnv := os.Getenv("APP_ENV")
	if strings.ToLower(appEnv) == "production" {
		EnvProduction()
	}else{
		EnvLocal()
	}
}
