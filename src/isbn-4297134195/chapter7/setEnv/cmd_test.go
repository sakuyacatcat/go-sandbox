package cmd_test

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateProfile(t *testing.T) {
	t.Setenv("DATABASE_URL", "localhost:5432")
	tEnv := os.Getenv("DATABASE_URL")
	fmt.Println(tEnv)
}
