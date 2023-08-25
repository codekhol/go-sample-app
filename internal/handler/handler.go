package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/codekhol/go-sample-app/internal/model"
	"github.com/codekhol/go-sample-app/internal/register"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	register register.Register
}

func New(reg register.Register) *Handler {
	return &Handler{
		register: reg,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading request body: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var user model.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		http.Error(w, fmt.Sprintf("error unmarshalling user: %s", err), http.StatusInternalServerError)
		return
	}

	hashed, err := hash(user.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("error hashing password: %s", err), http.StatusInternalServerError)
		return
	}
	user.Password = hashed

	err = h.register.Register(user)
	if err != nil {
		http.Error(w, fmt.Sprintf("error registering user: %s", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func hash(plain string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}
