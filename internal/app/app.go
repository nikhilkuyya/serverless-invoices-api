package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nikhilkuyya/invoice-go-app/internal/api"
	"github.com/nikhilkuyya/invoice-go-app/internal/store"
	"github.com/nikhilkuyya/invoice-go-app/migrations"
)


type Application struct {
	Logger *log.Logger
	BankAccountHandler *api.BankAccountHandler
	DB *sql.DB
}


func NewApplication() (*Application, error) {
	pgDB, err := store.Open()

	if err != nil {
		return nil, err
	}
	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}
	logger := log.New(os.Stdout,"invoice-app::",log.Ldate | log.Ltime)
	// stores
	bankAccountStore := store.NewPostgresBankAccountStore(pgDB)
	// handlers
	bankAccountHandler := api.NewBankAccountHandler(bankAccountStore)
	// app
	app := Application {
		Logger: logger,
		BankAccountHandler: bankAccountHandler,
		DB: pgDB,
	}
	return &app,nil
}


func (a *Application) HealthCheck(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"App health check is good\n")
}
