package main

import (
	"log"
	"os"
)

func main() {
	//conf := config.MustLoadConfig()

}

func setupLogger(env string) {
	switch env {
	case "local":
		log.SetOutput(os.Stdout)
	case "dev":
		log.SetOutput(os.Stderr)
	default:
		log.SetOutput(os.Stdout)
	}
}
