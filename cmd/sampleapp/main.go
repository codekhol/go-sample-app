package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/codekhol/go-sample-app/internal/handler"
	"github.com/codekhol/go-sample-app/internal/notifier"
	"github.com/codekhol/go-sample-app/internal/register"
	"github.com/codekhol/go-sample-app/internal/store"
	"github.com/gorilla/mux"
	_ "github.com/jackc/pgx/v5/stdlib"
)

const connStr = "postgres://postgres:postgres@db:5432/postgres?sslmode=disable"

func main() {
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		panic("error opening postgres connection: " + err.Error())
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		panic("error pinging postgres db")
	}

	store := store.New(db)
	notifier := notifier.New("smtp host", "smtp sender", "smtp password")
	register := register.New(store, notifier)
	handler := handler.New(register)
	router := mux.NewRouter()

	router.Handle("/register", handler).Methods(http.MethodPost)

	fmt.Println("starting server on localhost:8080...")
	log.Panic(http.ListenAndServe(":8080", router))
}
