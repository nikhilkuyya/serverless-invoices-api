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
	ClientHandler *api.ClientHandler
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
	clientStore := store.NewPostgresClientStore(pgDB)
	// handlers
	bankAccountHandler := api.NewBankAccountHandler(bankAccountStore)
	clientHandler := api.NewClientHandler(clientStore)
	// app
	app := Application {
		Logger: logger,
		BankAccountHandler: bankAccountHandler,
		ClientHandler: clientHandler,
		DB: pgDB,
	}
	return &app,nil
}


func (a *Application) HealthCheck(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"App health check is good\n")
}
