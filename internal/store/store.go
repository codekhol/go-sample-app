package store

import (
	"github.com/codekhol/go-sample-app/internal/model"
)

type Store interface {
	Save(user model.User) error
}
