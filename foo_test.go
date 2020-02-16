package main

import (
	"fmt"
	"os"
	"testing"
)

// TestFoo prints out the ENV.
func TestFoo(t *testing.T) {
	fmt.Println("Test foo is called.")
	env := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if env == "" {
		env = "default_value"
	}
	fmt.Println("FOO:", env[:5])
}
