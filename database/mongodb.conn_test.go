package database_test

import (
	database "mongodb-go/database"
	"testing"
)

func TestConnection(t *testing.T) {
	_, err := database.Connection()
	if err != nil {
		t.Errorf("Database connection failed")
	}
}
