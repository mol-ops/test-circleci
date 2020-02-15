package main

import (
	"fmt"
	"os"
	"testing"
)

func TestBar(t *testing.T) {
	fmt.Println("Test bar is called.")
	fmt.Println("BAR:", os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
}
