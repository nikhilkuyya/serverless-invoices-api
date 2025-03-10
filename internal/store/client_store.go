package store

import "database/sql"


type Client struct {
	ID int64 `json:"id"`
	CompanyName string `json:"company_name"`
	CompanyGSTNumber string `json:"company_gstin"`
	CompanyAddress string `json:"company_address"`
	CompanyCity string `json:"company_city"`
	CompanyPostalCode string `json:"company_postalcode"`
	CompanyState string `json:"company_state"`
	CompanyCountry string `json:"company_country"`
	CompanyEmail string `json:"company_email"`
	CompanyBankAccountId int64 `json:"company_bank_account_id"`
}

type ClientStore interface {
	CreateClient(client *Client) (*Client,error)
	GetClientByID(id int64) (*Client,error)
	GetClients() (*[]Client, error)
}

type PostgresClientStore struct {
	db *sql.DB
}

func NewPostgresClientStore(db *sql.DB) *PostgresClientStore{
	return &PostgresClientStore{
		db: db,
	}
}


func (pg *PostgresClientStore) CreateClient(client *Client) (*Client, error) {
	tx, err := pg.db.Begin()
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO clients
	(company_name, company_gstin, company_address, company_postalcode, company_city, company_state, company_country, company_email, company_bank_account_id)
	VALUES
	($1,$2, $3, $4, $5, $6, $7, $8, $9)
	RETURNING id
	`

	err = tx.QueryRow(query,client.CompanyName, client.CompanyGSTNumber, client.CompanyAddress, client.CompanyPostalCode, client.CompanyCity, client.CompanyState, client.CompanyCountry, client.CompanyEmail, client.CompanyBankAccountId).Scan(&client.ID)

	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return client,nil
}

func (pg *PostgresClientStore) GetClientByID(id int64) (*Client,error) {
	var client = Client{}
	const query = "SELECT id, company_name, company_gstin, company_address, company_postalcode, company_city, company_state, company_country, company_email from clients WHERE id = $1"

	 err := pg.db.QueryRow(query,id).Scan(&client.ID, &client.CompanyName, &client.CompanyGSTNumber,&client.CompanyAddress,  &client.CompanyPostalCode, &client.CompanyCity, &client.CompanyState, &client.CompanyCountry, &client.CompanyEmail)

	 if err == sql.ErrNoRows {
		return nil, nil
	}
	 if err != nil {
		return nil, err
	 }

	 return &client,nil
}

func (pg *PostgresClientStore) GetClients() (*[]Client, error) {
	const query = "SELECT id, company_name, company_gstin, company_address, company_postalcode, company_city, company_state, company_country, company_email FROM clients LIMIT 100"
	rows, err := pg.db.Query(query);

	if err != nil {
		return nil, err
	}
	defer rows.Close();

	var clients []Client
	for rows.Next() {
		var client = Client{}
		err = rows.Scan(&client.ID, &client.CompanyName, &client.CompanyGSTNumber, &client.CompanyAddress, &client.CompanyPostalCode, &client.CompanyCity, &client.CompanyState, &client.CompanyCountry, &client.CompanyEmail)

		if err != nil {
			return nil, err
		}
		clients = append(clients, client);
	}
	return &clients, nil

}
