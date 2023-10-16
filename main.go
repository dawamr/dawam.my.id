package main

import (
	"fmt"
	"log"

	"github.com/dawamr/resume-cv-2023/bootstrap"
	"github.com/dawamr/resume-cv-2023/pkg/env"
)

func main() {
	app := bootstrap.NewApplication()
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", env.GetEnv("APP_HOST", "127.0.0.1"), env.GetEnv("APP_PORT", "4000"))))
}
