package store_test

import (
	"database/sql"
	"fmt"
	"math/rand"
	"testing"

	"github.com/codekhol/go-sample-app/internal/model"
	"github.com/codekhol/go-sample-app/internal/store"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
)

type testdb struct {
	name string
	db   *sql.DB
}

func TestSave(t *testing.T) {
	testDB := setup(t)
	defer teardown(t, testDB)
	s := store.New(testDB.db)

	testTable := []struct {
		name      string
		user      model.User
		expectErr bool
	}{
		{
			name:      "success 1",
			user:      model.User{Email: "email1@mail.com", Password: "123456"},
			expectErr: false,
		},
		{
			name:      "success 2",
			user:      model.User{Email: "email2@mail.com", Password: "123456"},
			expectErr: false,
		},
		{
			name:      "failed. duplicate email",
			user:      model.User{Email: "email1@mail.com", Password: "123456"},
			expectErr: true,
		},
		{
			name:      "failed. empty email",
			user:      model.User{Email: "", Password: "123456"},
			expectErr: true,
		},
		{
			name:      "failed. empty password",
			user:      model.User{Email: "email1@mail.com", Password: ""},
			expectErr: true,
		},
	}

	for _, test := range testTable {
		err := s.Save(test.user)
		if (err == nil) == test.expectErr {
			t.Errorf("error in test `%s`. expected error %t, got %s", test.name, test.expectErr, err)
		}
	}
}

func setup(t *testing.T) testdb {
	conn, err := sql.Open("pgx", connStr("postgres"))
	if err != nil {
		t.Fatalf("error opening postgres connection: %s", err.Error())
	}
	defer conn.Close()

	testDBName := fmt.Sprintf("testdb_%d", rand.Int())
	_, err = conn.Exec(fmt.Sprintf("CREATE DATABASE %s", testDBName))
	if err != nil {
		t.Fatalf("error creating test database: %s", err.Error())
	}

	testDB, err := sql.Open("pgx", connStr(testDBName))
	if err != nil {
		t.Fatalf("error opening test connection: %s", err.Error())
	}

	mig, err := migrate.New("file://../../migrations", connStr(testDBName))
	if err != nil {
		t.Fatalf("error creating migrations: %s", err.Error())
	}

	err = mig.Up()
	if err != nil {
		t.Fatalf("error running migrations: %s", err.Error())
	}

	return testdb{name: testDBName, db: testDB}
}

func teardown(t *testing.T, testDB testdb) {
	db, err := sql.Open("pgx", connStr("postgres"))
	if err != nil {
		t.Fatalf("error opening postgres connection: %s", err.Error())
	}
	defer db.Close()

	err = testDB.db.Close()
	if err != nil {
		t.Fatalf("error closing test connection: %s", err.Error())
	}

	_, err = db.Exec(fmt.Sprintf("DROP DATABASE %s WITH (force)", testDB.name))
	if err != nil {
		t.Fatalf("error dropping test database %s: %s", testDB.name, err.Error())
	}
}

func connStr(dbname string) string {
	return fmt.Sprintf("postgres://postgres:postgres@db:5432/%s?sslmode=disable", dbname)
}
