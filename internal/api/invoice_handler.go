package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nikhilkuyya/invoice-go-app/internal/store"
)

type CreateInvoice struct {
	ID int `json:"id"`
	Number string `json:"number"`
	StatusId int `json:"status_id"`
	TeamId int64 `json:"team_id"`
	BankAccountId int64 `json:"bank_account_id"`
	ClientId	int64 `json:"client_id"`
	ConsigneeId int `json:"consignee_id"`
	Notes string `json:"notes"`
	Total int64 `json:"total"`
	Rows []CreateInvoiceRow `json:"rows"`
}

type CreateInvoiceRow struct {
	ID int64 `json:"id"`
	SerialNumber int64 `json:"serial_no"`
	Item string `json:"item"`
	Description string `json:"description"`
	HSNCode string `json:"hsn_code"`
	Quantity int64 `json:"quantity"`
	Price int64 `json:"price"`
	Unit string `json:"unit"`
	InvoiceRowOrder int `json:"invoice_row_order"`
	InvoiceTaxes []int `json:"invoice_taxes"`
}

type InvoiceHandler struct {
	teamStore store.TeamStore
	clientStore store.ClientStore
	taxStore store.TaxStore
	invoiceStore store.InvoiceStore
}

func NewInvoiceHandler(teamStore store.TeamStore, clientStore store.ClientStore, taxStore store.TaxStore, invoiceStore store.InvoiceStore) *InvoiceHandler {
	return &InvoiceHandler{
		teamStore: teamStore,
		clientStore: clientStore,
		taxStore: taxStore,
		invoiceStore: invoiceStore,
	}
}

func (invoiceHandler *InvoiceHandler) HandleCreateInovice(w http.ResponseWriter, r *http.Request) {
	var createInvoice = CreateInvoice{}
	json.NewDecoder(r.Body).Decode(&createInvoice);
	fmt.Println(createInvoice);
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(createInvoice)
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
