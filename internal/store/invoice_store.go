package store

import "database/sql"


type InvoiceTax struct {
	Id int64 `json:"id"`
	Name int64 `json:"name"`
	Label string `json:"label"`
	TaxPercentage int64 `json:"tax_percentage"`
}

type InvoiceRow struct {
	ID int64 `json:"id"`
	SerialNumber int64 `json:"serial_no"`
	Item string `json:"item"`
	Description string `json:"description"`
	HSNCode string `json:"hsn_code"`
	Quantity int64 `json:"quantity"`
	Price int64 `json:"price"`
	Unit string `json:"unit"`
	InvoiceRowOrder int `json:"invoice_row_order"`
	InvoiceTaxes []InvoiceTax `json:"invoice_taxes"`
}

type InvoiceStatus struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Label string `json:"label"`
}

type Invoice struct {
	ID int64 `json:"id"`
	Number string `json:"number"`
	IssuedAt string `json:"issued_at"`
	DueAt string `json:"due_at"`
	Currency string `json:"currency"`
	FromName string `json:"from_name"`
	FromGSTNumber string `json:"from_gstin"`
	FromAddress string `json:"from_address"`
	FromCity string `json:"from_city"`
	FromState string `json:"from_state"`
	FromCountry string `json:"from_country"`
	FromEmail string `json:"from_email"`
	FromPhone string `json:"from_phone"`
	TeamId int64 `json:"team_id"`
	BankName string `json:"bank_name"`
	BankAccountNumber string `json:"bank_account_number"`
	BankIFSCCode string `json:"bank_ifsc_code"`
	ClientName string `json:"client_name"`
	ClientGstNumber string `json:"client_gstin"`
	ClientAddress string `json:"client_address"`
	ClientPostalCode string `json:"client_postal_code"`
	ClientCity string `json:"client_city"`
	ClientState string `json:"client_state"`
	ClientCountry string `json:"client_country"`
	ClientId string `json:"client_id"`
	ConsigneeName string `json:"consignee_name"`
	ConsigneeGstNumber string `json:"consignee_gstin"`
	ConsigneeAddress string `json:"consignee_adress"`
	ConsigneePostalCode string `json:"consignee_postal_code"`
	ConsigneeCity string `json:"consignee_city"`
	ConsigneeState string `json:"consignee_state"`
	ConsigneeCountry string `json:"consignee_country"`
	ConsigneeEmail string `json:"consignee_email"`
	ConsigneeId int64 `json:"consignee_id"`
	Notes string `json:"notes"`
	Status InvoiceStatus `json:"status"`
	Rows []InvoiceRow `json:"rows"`
	Total int64 `json:"total"`
}

type InvoiceStore interface {
	CreateInvoice(invoice *Invoice) (*Invoice, error)
	GetInvoiceByID(id int64) (*Invoice, error)
	GetInvoices() (*[]Invoice, error)
	GetInvoiceStatuses() (*[]InvoiceStatus, error)
}

type PostgresInvoiceStore struct {
	db *sql.DB
}

func NewPostgresInvoiceStore(db *sql.DB) *PostgresInvoiceStore {
	return &PostgresInvoiceStore{
		db: db,
	}
}

func (pg *PostgresInvoiceStore) CreateInvoice(invoice *Invoice) (*Invoice, error) {
	return nil,nil
}

func (pg *PostgresInvoiceStore) GetInvoiceByID(id int64) (*Invoice, error) {
	return nil,nil
}

func (pg *PostgresInvoiceStore) GetInvoices() (*[]Invoice, error) {
	return nil,nil
}

func (pg *PostgresInvoiceStore) GetInvoiceStatuses() (*[]InvoiceStatus, error) {
	const query = `SELECT id, name, label FROM invoice_statuses LIMIT 10`
	rows, err :=pg.db.Query(query)
	if err != nil {
		return nil, err
	}
	var invoiceStatuses []InvoiceStatus = make([]InvoiceStatus, 0, 10)
	for rows.Next() {
		var invoiceStatus = InvoiceStatus{}
		err = rows.Scan(&invoiceStatus.Id, &invoiceStatus.Name, &invoiceStatus.Label)
		if err != nil {
			return nil, err
		}
		invoiceStatuses = append(invoiceStatuses, invoiceStatus)
	}
	return &invoiceStatuses, nil
}
