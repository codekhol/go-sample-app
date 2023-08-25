package register_test

import (
	"errors"
	"testing"

	"github.com/codekhol/go-sample-app/internal/model"
	"github.com/codekhol/go-sample-app/internal/notifier"
	"github.com/codekhol/go-sample-app/internal/register"
)

func TestRegister(t *testing.T) {
	testTable := []struct {
		name      string
		user      model.User
		expectErr bool
	}{
		{
			name:      "success",
			user:      model.User{Email: "test@test.com", Password: "1234567890"},
			expectErr: false,
		},
		{
			name:      "store fails",
			user:      model.User{Email: "duplicate", Password: "1234567890"},
			expectErr: true,
		},
		{
			name:      "notifier fails",
			user:      model.User{Email: "invalid", Password: "1234567890"},
			expectErr: true,
		},
	}

	s := &mockStore{}
	n := &mockNotifier{}
	r := register.New(s, n)

	for _, test := range testTable {
		err := r.Register(test.user)
		if (err == nil) == test.expectErr {
			t.Errorf("test `%s` failed. expected error %t, got %s", test.name, test.expectErr, err)
		}
	}
}

type mockNotifier struct {
}

func (m *mockNotifier) Send(n notifier.Notification) error {
	if n.Destination == "invalid" {
		return errors.New("unable to notify: invalid destination")
	}

	// NOOP

	return nil
}

type mockStore struct {
}

func (m *mockStore) Save(user model.User) error {
	if user.Email == "duplicate" {
		return errors.New("error saving user: email already exists")
	}

	// NOOP

	return nil
}
