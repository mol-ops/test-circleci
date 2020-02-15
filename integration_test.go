// +build integration

package main

import (
	"fmt"
	"os"
	"testing"
)

func TestIntegration(t *testing.T) {
	fmt.Println("Test integration is called.")
	fmt.Println("INT:", os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
}
