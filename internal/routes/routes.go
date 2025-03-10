package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nikhilkuyya/invoice-go-app/internal/app"
)

func SetupRoutes(app *app.Application) (*chi.Mux){
	r := chi.NewRouter()
	r.Get("/health", app.HealthCheck)
	r.Get("/bank-account/{id}",app.BankAccountHandler.HandleGetBankAccountByID)
	r.Post("/bank-account",app.BankAccountHandler.HandleCreateBankAccount)
	r.Get("/bank-account/list",app.BankAccountHandler.HandleGetAllBankAccounts);

	r.Get("/client/{id}",app.ClientHandler.HandleGetClientByID)
	r.Post("/client",app.ClientHandler.HandleCreateClient)
	r.Get("/client/list",app.ClientHandler.HandleGetClients);
	return r
}
