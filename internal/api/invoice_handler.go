package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/nikhilkuyya/invoice-go-app/internal/store"
)

type CreateInvoice struct {
	ID int `json:"id"`
	Number string `json:"invoice_number"`
	IssuedAt string `json:"issued_at"`
	StatusId int64 `json:"status_id"`
	TeamId int64 `json:"team_id"`
	BankAccountId int64 `json:"bank_account_id"`
	ClientId	int64 `json:"client_id"`
	ConsigneeId int64 `json:"consignee_id"`
	Notes string `json:"notes"`
	Total int64 `json:"total"`
	Rows []CreateInvoiceRow `json:"rows"`
}

type CreateInvoiceRow struct {
	ID int64 `json:"id"`
	SerialNumber string `json:"serial_no"`
	Item string `json:"item"`
	Description string `json:"description"`
	HSNCode string `json:"hsn_code"`
	Quantity int64 `json:"quantity"`
	Price int64 `json:"price"`
	Unit string `json:"unit"`
	InvoiceRowOrder int `json:"invoice_row_order"`
}

type InvoiceHandler struct {
	teamStore store.TeamStore
	clientStore store.ClientStore
	taxStore store.TaxStore
	invoiceStore store.InvoiceStore
	bankAccountStore store.BankAccountStore
}

func NewInvoiceHandler(teamStore store.TeamStore, clientStore store.ClientStore, taxStore store.TaxStore, invoiceStore store.InvoiceStore, bankAccountStore store.BankAccountStore) *InvoiceHandler {
	return &InvoiceHandler{
		teamStore: teamStore,
		clientStore: clientStore,
		taxStore: taxStore,
		invoiceStore: invoiceStore,
		bankAccountStore: bankAccountStore,
	}
}

func (invoiceHandler *InvoiceHandler) HandleCreateInovice(w http.ResponseWriter, r *http.Request) {
	var createInvoice = CreateInvoice{}
	json.NewDecoder(r.Body).Decode(&createInvoice);
	team, err := invoiceHandler.teamStore.GetTeamByID(createInvoice.TeamId)
	if err != nil {
		fmt.Println(err)
		http.Error(w,"failed to created invoice",http.StatusBadRequest)
		return
	}

	client, err := invoiceHandler.clientStore.GetClientByID(createInvoice.ClientId)
	if err != nil {
		fmt.Println(err)
		http.Error(w,"failed to create invoice",http.StatusBadRequest)
		return
	}

	consignee, err := invoiceHandler.clientStore.GetClientByID(createInvoice.ConsigneeId)
	if err != nil {
		fmt.Println(err)
		http.Error(w,"failed to create invoice",http.StatusBadRequest)
		return
	}

	bankAccount, err := invoiceHandler.bankAccountStore.GetBankAccountByID(createInvoice.BankAccountId)
	if err != nil {
		fmt.Println(err)
		http.Error(w,"failed to create invoice",http.StatusBadRequest)
		return
	}

	taxes, err := invoiceHandler.taxStore.GetTaxes()
	if err != nil {
		fmt.Println(err)
		http.Error(w,"failed to create invoice",http.StatusBadRequest)
		return;
	}

	issuedAtDate,err := time.Parse("2006-01-02",createInvoice.IssuedAt)
	if err != nil {
		fmt.Println(err)
		http.Error(w,"faild to create invoice",http.StatusBadRequest)
		return
	}

	dueAt := issuedAtDate.AddDate(0,0,team.InvoiceDueDays).Format("2006-01-02")
	invoiceRows := make([]store.InvoiceRow,0);
	invoiceTaxes := make([]store.InvoiceTax,0);
	for _, taxInfo := range *taxes{
		var invoiceTax = store.InvoiceTax{
			Name: taxInfo.Name,
			Label: taxInfo.Label,
			TaxPercentage: taxInfo.TaxPercentage,
		}
		invoiceTaxes = append(invoiceTaxes, invoiceTax);
	}

	for _, iterationRow := range createInvoice.Rows {
		var invoiceRow = store.InvoiceRow{
			SerialNumber: iterationRow.SerialNumber,
			Item: iterationRow.Item,
			Description: iterationRow.Description,
			HSNCode: iterationRow.HSNCode,
			Quantity: iterationRow.Quantity,
			Price: iterationRow.Price,
			Unit: iterationRow.Unit,
			InvoiceRowOrder: iterationRow.InvoiceRowOrder,
			InvoiceTaxes: &invoiceTaxes,
		}
		invoiceRows = append(invoiceRows, invoiceRow)
	}

	createInvoicePayload := store.Invoice{
		Number: createInvoice.Number,
		IssuedAt: createInvoice.IssuedAt,
		Status: &store.InvoiceStatus {
			Id: int64(createInvoice.StatusId),
		},
		DueAt: dueAt,
		Currency: team.Currency,
		FromName: team.CompanyName,
		FromGSTNumber: team.CompanyGSTNumber,
		FromAddress: team.CompanyAddress,
		FromCity: team.CompanyCity,
		FromState: team.CompanyState,
		FromCountry: team.CompanyCountry,
		FromEmail: team.ContactEmail,
		FromPhone: team.ContactPhone,
		FromPostalCode: team.CompanyPostalCode,
		TeamId: team.ID,
		BankName: bankAccount.BankName,
		BankAccountNumber: bankAccount.BankAccountNumber,
		BankIFSCCode: bankAccount.BankIfscCode,
		ClientName: client.CompanyName,
		ClientGstNumber: client.CompanyGSTNumber,
		ClientAddress: client.CompanyAddress,
		ClientPostalCode: client.CompanyPostalCode,
		ClientCity: client.CompanyCity,
		ClientState: client.CompanyState,
		ClientCountry: client.CompanyCountry,
		ClientEmail: client.CompanyEmail,
		ClientId: client.ID,
		ConsigneeName: consignee.CompanyName,
		ConsigneeGstNumber: consignee.CompanyGSTNumber,
		ConsigneeAddress: consignee.CompanyAddress,
		ConsigneePostalCode: consignee.CompanyPostalCode,
		ConsigneeCity: consignee.CompanyCity,
		ConsigneeState: consignee.CompanyState,
		ConsigneeCountry: consignee.CompanyCountry,
		ConsigneeEmail: consignee.CompanyEmail,
		ConsigneeId: consignee.ID,
		Notes: createInvoice.Notes,
		Total: createInvoice.Total,
		Rows: &invoiceRows,
	}

	response,err := invoiceHandler.invoiceStore.CreateInvoice(&createInvoicePayload)
	if err != nil {
		fmt.Println(err)
		http.Error(w,"failed to create invoice",http.StatusInternalServerError);
		return
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(response)
}

func (invoiceHandler *InvoiceHandler) HandleGetInvoice(w http.ResponseWriter, r *http.Request) {
}

func (invoiceHandler *InvoiceHandler) HandleInvoices(w http.ResponseWriter, r *http.Request) {
}

func (invoiceHandler *InvoiceHandler) HandleInvoiceStatues(w http.ResponseWriter, r *http.Request){
	invoiceStatus, err := invoiceHandler.invoiceStore.GetInvoiceStatuses()
	if err != nil {
		fmt.Println(err)
		http.Error(w,"Failed to fetch invoice status",http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(invoiceStatus)
}
