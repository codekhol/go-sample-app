package store

import (
	"database/sql"

	"github.com/codekhol/go-sample-app/internal/model"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pkg/errors"
)

type pgStore struct {
	db *sql.DB
}

func New(db *sql.DB) Store {
	return &pgStore{
		db: db,
	}
}

func (pg *pgStore) Save(user model.User) error {
	err := validate(user)
	if err != nil {
		return errors.Wrap(err, "error validating user")
	}

	_, err = pg.db.Exec("INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)
	if err != nil {
		return errors.Wrap(err, "error storing user")
	}

	return nil
}

func validate(user model.User) error {
	if user.Email == "" {
		return errors.New("email cannot be empty")
	}

	if user.Password == "" {
		return errors.New("password cannot be empty")
	}

	return nil
}
