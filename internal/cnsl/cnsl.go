package cnsl

import (
	"fmt"
	"log"

	"github.com/ad9311/renio-go/internal/conf"
)

func AppName() {
	fmt.Print("\n--- RENIO ---\n\n")
}

func Success(message string) {
	if conf.GetEnv().AppEnv != "test" {
		fmt.Printf("✓ %s\n", message)
	}
}

func Info(message string) {
	if conf.GetEnv().AppEnv != "test" {
		fmt.Printf("✓ %s\n", message)
	}
}

func Error(message string) {
	if conf.GetEnv().AppEnv != "test" {
		fmt.Printf("x %s\n", message)
	}
}

func Fatal(message string) {
	log.Fatalf("FATAL: %s\n", message)
}
