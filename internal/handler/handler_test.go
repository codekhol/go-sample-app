package handler_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codekhol/go-sample-app/internal/handler"
	"github.com/codekhol/go-sample-app/internal/model"
)

func TestServeHTTP(t *testing.T) {
	testTable := []struct {
		name         string
		req          *http.Request
		expectedCode int
	}{
		{
			name:         "success",
			req:          makeRequest(&model.User{Email: "test@test.com", Password: "1234567890"}, t),
			expectedCode: 201,
		},
		{
			name:         "register fails",
			req:          makeRequest(&model.User{Email: "invalid", Password: "1234567890"}, t),
			expectedCode: 500,
		},
		{
			name:         "unmarshal fails",
			req:          makeRequest([]string{"invalid", "body"}, t),
			expectedCode: 500,
		},
	}

	r := &mockRegister{}
	h := handler.New(r)

	for _, test := range testTable {
		rr := httptest.ResponseRecorder{}
		h.ServeHTTP(&rr, test.req)
		if rcvdCode := rr.Result().StatusCode; rcvdCode != test.expectedCode {
			t.Errorf("test `%s` failed. expected status code: %d got %d", test.name, test.expectedCode, rcvdCode)
		}
	}
}

func makeRequest(obj interface{}, t *testing.T) *http.Request {
	data, err := json.Marshal(obj)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, "dummy", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	return req
}

type mockRegister struct {
}

func (s *mockRegister) Register(user model.User) error {
	if user.Email == "invalid" {
		return errors.New("error registering user")
	}

	return nil
}
