package main

import (
	"fmt"
	"log"

	"github.com/kooroshh/fiber-boostrap/bootstrap"
	"github.com/kooroshh/fiber-boostrap/pkg/env"
)

func main() {
	app := bootstrap.NewApplication()
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", env.GetEnv("APP_HOST", "0.0.0.0"), env.GetEnv("APP_PORT", "4000"))))
}