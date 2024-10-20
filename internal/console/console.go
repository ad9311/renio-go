package console

import (
	"fmt"
	"log"
	"os"
)

var env = os.Getenv("RENIO_ENV")

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
	log.Fatalf("x %s\n", message)
}

func Query(query string) {
	if env != "test" {
		fmt.Printf("BEGIN `%s`\n", query)
	}
}
