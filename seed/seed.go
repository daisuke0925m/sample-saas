package db

import (
	"database/sql"
	"testing"

	"github.com/tanimutomo/sqlfile"
)

func Seed(t *testing.T, conn *sql.DB) {
	s := sqlfile.New()
	if err := s.File("../db/seed.go"); err != nil {
		t.Fatalf("failed to open DBMS:%v", err)
	}

	_, err := s.Exec(conn)
	if err != nil {
		t.Fatalf("failed to seed: %v", err)
	}
}
