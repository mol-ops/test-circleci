package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFoo(t *testing.T) {
	fmt.Println("Test foo is called.")
	fmt.Println("FOO:", os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
}
