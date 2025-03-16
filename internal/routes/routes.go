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

	r.Get("/team/{id}",app.TeamHandler.HandleGetTeamByID)
	r.Post("/team",app.TeamHandler.HandleCreateTeam)
	r.Get("/team",app.TeamHandler.HandleGetTeams)

	r.Get("/tax/{id}",app.TaxHandler.HandleGetTaxByID)
	r.Get("/tax", app.TaxHandler.HandleGetTaxes)
	r.Post("/tax",app.TaxHandler.HandleCreateTax)

	r.Post("/invoice",app.InvoiceHandler.HandleCreateInovice)
	r.Get("/invoice-status", app.InvoiceHandler.HandleInvoiceStatues)
	return r
}
