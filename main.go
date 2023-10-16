package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dawamr/resume-cv-2023/bootstrap"
	"github.com/dawamr/resume-cv-2023/pkg/env"
)

func main() {
	app := bootstrap.NewApplication()
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", env.GetEnv("APP_HOST", os.Getenv("APP_HOST")), env.GetEnv("APP_PORT", os.Getenv("APP_PORT")))))
}
