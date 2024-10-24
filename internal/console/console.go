package console

import (
	"fmt"
	"log"

	"github.com/ad9311/renio-go/internal/envs"
)

var env = envs.GetEnvs().ENV

func AppName() {
	fmt.Print("\n--- RENIO ---\n\n")
}

func Success(message string) {
	if env != "test" {
		fmt.Printf("✓ %s\n", message)
	}
}

func Info(message string) {
	if env != "test" {
		fmt.Printf("✓ %s\n", message)
	}
}

func Error(message string) {
	if env != "test" {
		fmt.Printf("x %s\n", message)
	}
}

func Fatal(message string) {
	log.Fatalf("FATAL: %s\n", message)
}
