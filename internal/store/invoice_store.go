package store

import (
	"database/sql"
	"fmt"
	"strings"
)


type InvoiceTax struct {
	Id int64 `json:"id"`
	InvoiceRowId int64 `json:"invoice_row_id"`
	Name string `json:"name"`
	Label string `json:"label"`
	TaxPercentage int64 `json:"tax_percentage"`
}

type InvoiceRow struct {
	ID int64 `json:"id"`
	SerialNumber string `json:"serial_no"`
	Item string `json:"item"`
	Description string `json:"description"`
	HSNCode string `json:"hsn_code"`
	Quantity int64 `json:"quantity"`
	Price int64 `json:"price"`
	Unit string `json:"unit"`
	InvoiceRowOrder int `json:"invoice_row_order"`
	InvoiceTaxes *[]InvoiceTax `json:"invoice_taxes"`
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
	FromPostalCode string `json:"from_postal_code"`
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
	ClientEmail string `json:"client_email"`
	ClientId int64 `json:"client_id"`
	ConsigneeName string `json:"consignee_name"`
	ConsigneeGstNumber string `json:"consignee_gstin"`
	ConsigneeAddress string `json:"consignee_address"`
	ConsigneePostalCode string `json:"consignee_postal_code"`
	ConsigneeCity string `json:"consignee_city"`
	ConsigneeState string `json:"consignee_state"`
	ConsigneeCountry string `json:"consignee_country"`
	ConsigneeEmail string `json:"consignee_email"`
	ConsigneeId int64 `json:"consignee_id"`
	Notes string `json:"notes"`
	Status *InvoiceStatus `json:"status"`
	Rows *[]InvoiceRow `json:"rows"`
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

// Function to generate placeholders dynamically
func generatePlaceholders(count int) string {
	placeholders := make([]string, count)
	for i := range placeholders {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}
	return strings.Join(placeholders, ", ")
}

func (pg *PostgresInvoiceStore) CreateInvoice(invoice *Invoice) (*Invoice, error) {
	tx, err := pg.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	invoiceColumnAndData := map[string]interface{}{
		"number": invoice.Number,
		"status_id": invoice.Status.Id,
		"issued_at": invoice.IssuedAt,
		"due_at": invoice.DueAt,
		"notes": invoice.Notes,
		"total": invoice.Total,
		"currency": invoice.Currency,
		"team_id": invoice.TeamId,
		"from_name": invoice.FromName,
		"from_gstin": invoice.FromGSTNumber,
		"from_address": invoice.FromAddress,
		"from_postal_code": invoice.FromPostalCode,
		"from_city": invoice.FromCity,
		"from_state": invoice.FromState,
		"from_email": invoice.FromEmail,
		"from_phone": invoice.FromPhone,
		"from_country": invoice.FromCountry,
		"bank_name": invoice.BankName,
		"bank_account_number": invoice.BankAccountNumber,
		"bank_ifsc_code": invoice.BankIFSCCode,
		"client_name": invoice.ClientName,
		"client_gstin": invoice.ClientGstNumber,
		"client_address": invoice.ClientAddress,
		"client_postal_code": invoice.ClientPostalCode,
		"client_city": invoice.ClientCity,
		"client_state": invoice.ClientState,
		"client_country": invoice.ClientCountry,
		"client_email": invoice.ClientEmail,
		"client_id":  invoice.ClientId,
		"consignee_name": invoice.ConsigneeName,
		"consignee_gstin": invoice.ConsigneeGstNumber,
		"consignee_address": invoice.ConsigneeAddress,
		"consignee_postal_code": invoice.ConsigneePostalCode,
		"consignee_city": invoice.ConsigneeCity,
		"consignee_state": invoice.ConsigneeState,
		"consignee_country": invoice.ConsigneeCountry,
		"consignee_email": invoice.ConsigneeEmail,
		"consignee_id": invoice.ConsigneeId,
	}

	// Extract column names and values
	columns := make([]string, 0, len(invoiceColumnAndData))
	values := make([]interface{}, 0, len(invoiceColumnAndData))
	// placeholders := make([]string, 0, len(invoiceColumnAndData))

	for col, val := range invoiceColumnAndData {
		columns = append(columns, col)
		values = append(values, val)
	}
	// placeholders = generatePlaceholders(len(values))

	// Construct query dynamically
	query := fmt.Sprintf("INSERT INTO invoices (%s) VALUES (%s) RETURNING id",
		strings.Join(columns, ", "), generatePlaceholders(len(values)))

	err = tx.QueryRow(query, values...).Scan(&invoice.ID)

	if err != nil {
		return nil, err
	}

	for _,invoiceRow := range *invoice.Rows {
		query := `INSERT INTO invoice_rows (invoice_id, serial_no, item, description, hsn_code, quantity, price, unit, invoice_row_order) VALUES ($1, $2, $3,$4, $5, $6, $7, $8, $9
		) RETURNING id`
		err = tx.QueryRow(query,invoice.ID, invoiceRow.SerialNumber, invoiceRow.Item, invoiceRow.Description, invoiceRow.HSNCode, invoiceRow.Quantity, invoiceRow.Price, invoiceRow.Unit, invoiceRow.InvoiceRowOrder).Scan(&invoiceRow.ID)
		if err != nil {
			return nil,err
		}

		for _, invoiceTax := range *invoiceRow.InvoiceTaxes {
			query := `INSERT INTO invoice_taxes
			(invoice_row_id, name, label, tax_percentage) VALUES($1, $2, $3, $4) RETURNING id`
			err = tx.QueryRow(query,invoiceRow.ID, invoiceTax.Name, invoiceTax.Label, invoiceTax.TaxPercentage).Scan(&invoiceTax.Id)
			if err != nil {
				return nil, err
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return invoice,nil
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
