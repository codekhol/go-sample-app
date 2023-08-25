package register

import (
	"fmt"

	"github.com/codekhol/go-sample-app/internal/model"
	"github.com/codekhol/go-sample-app/internal/notifier"
	"github.com/codekhol/go-sample-app/internal/store"
	"github.com/pkg/errors"
)

type userRegister struct {
	store   store.Store
	notifer notifier.Notifier
}

func New(store store.Store, notifier notifier.Notifier) Register {
	return &userRegister{
		store:   store,
		notifer: notifier,
	}
}

func (r *userRegister) Register(user model.User) error {
	err := r.store.Save(user)
	if err != nil {
		return errors.Wrap(err, "unable to store user")
	}

	notification := notifier.Notification{
		Destination: user.Email,
		Content:     fmt.Sprintf("Welcome %s! Thank you for registering with us.", user.Email),
	}
	err = r.notifer.Send(notification)
	if err != nil {
		return errors.Wrap(err, "unable to notify user")
	}

	return nil
}
