package store

import "database/sql"

type Team struct {
	ID int64 `json:"id"`
	CompanyName string `json:"company_name"`
	CompanyGSTNumber string `json:"company_gstin"`
	CompanyAddress string `json:"company_address"`
	CompanyPostalCode string `json:"company_postal_code"`
	CompanyCity string `json:"company_city"`
	CompanyState string `json:"company_state"`
	CompanyCountry string `json:"company_country"`
	Website string `json:"website"`
	ContactEmail string `json:"contact_email"`
	ContactPhone string `json:"contact_phone"`
	Currency string `json:"currency"`
	InvoiceDueDays int64 `json:"invoice_due_days"`
	InvoiceLateFee int64 `json:"invoice_late_fee"`
	LogoUrl string `json:"logo_url"`
}

type TeamStore interface {
	CreateTeam(team *Team) (*Team, error)
	GetTeamByID(id int64) (*Team,error)
	GetTeams() (*[]Team,error)
}

type PostgresTeamStore struct {
	db *sql.DB
}

func NewPostgresTeamStore(db *sql.DB) *PostgresTeamStore {
	return &PostgresTeamStore {
		db: db,
	}
}

func (pg *PostgresTeamStore) CreateTeam(team *Team) (*Team, error){
	tx, err := pg.db.Begin()
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO teams
	(company_name, company_gstin, company_address, company_postal_code, company_city, company_state, company_country, website, contact_email, contact_phone, currency, invoice_due_days, invoice_late_fee, logo_url) VALUES
	($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
	RETURNING id`

	err = tx.QueryRow(query, team.CompanyName, team.CompanyGSTNumber, team.CompanyAddress, team.CompanyPostalCode, team.CompanyCity, team.CompanyState, team.CompanyCountry, team.Website, team.ContactEmail, team.ContactPhone, team.Currency, team.InvoiceDueDays, team.InvoiceDueDays, team.LogoUrl).Scan(&team.ID)

	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return team, nil
}

func (pg *PostgresTeamStore) GetTeamByID(teamID int64) (*Team, error) {
	var team = Team{}
	const query = `SELECT id, company_name, company_gstin, company_address, company_postal_code, company_city, company_state, company_country, website, contact_email, contact_phone, currency, invoice_due_days, invoice_late_fee, logo_url FROM teams WHERE id = $1`

	err := pg.db.QueryRow(query,teamID).Scan(
		&team.ID, &team.CompanyName, &team.CompanyGSTNumber, &team.CompanyAddress, &team.CompanyPostalCode, &team.CompanyCity, &team.CompanyState, &team.CompanyCountry, &team.Website, &team.ContactEmail, &team.ContactPhone, &team.Currency, &team.InvoiceDueDays, &team.InvoiceLateFee, &team.LogoUrl)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &team,nil
}


func (pg *PostgresTeamStore) GetTeams() (*[]Team, error) {
	const query = `SELECT id, company_name, company_gstin, company_address, company_postal_code, company_city, company_state, company_country, website, contact_email, contact_phone, currency, invoice_due_days, invoice_late_fee, logo_url FROM teams LIMIT 100`

	rows, err := pg.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var teams []Team
	for rows.Next() {
		var team = Team{}

		err = rows.Scan(
			&team.ID, &team.CompanyName, &team.CompanyGSTNumber, &team.CompanyAddress, &team.CompanyPostalCode, &team.CompanyCity, &team.CompanyState, &team.CompanyCountry, &team.Website, &team.ContactEmail, &team.ContactPhone, &team.Currency, &team.InvoiceDueDays, &team.InvoiceLateFee, &team.LogoUrl)


		if err != nil {
			return nil, err
		}

		teams = append(teams,team)
	}

	return &teams, nil
}
