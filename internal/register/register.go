package register

import (
	"github.com/codekhol/go-sample-app/internal/model"
)

type Register interface {
	Register(user model.User) error
}
