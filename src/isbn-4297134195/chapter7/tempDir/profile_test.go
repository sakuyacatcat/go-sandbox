package profile_test

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateProfile(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "profile.json")
	got, err := CreateProfile(filename)
	if err != nil {
		t.Fatal(err)
	}
	want := true
	if got != want {
		t.Fatalf("want %v, but %v:", want, got)
	}
}

func CreateProfile(filename string) (bool, error) {
	file, err := os.Create(filename)
	if err != nil {
		return false, err
	}
	defer file.Close()

	_, err = file.WriteString("Sample data for profile")
	if err != nil {
		return false, err
	}

	return true, nil
}
