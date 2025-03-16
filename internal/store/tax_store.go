package store

import "database/sql"

type Tax struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Label string `json:"label"`
	TaxPercentage int64 `json:"tax_percentage"`
}

type TaxStore interface {
	CreateTax(tax *Tax) (*Tax, error)
	GetTaxByID(id int64) (*Tax, error)
	GetTaxes() (*[]Tax, error)
}


type PostgresTaxStore struct {
	db *sql.DB
}

func NewPostgresTaxStore(db *sql.DB) *PostgresTaxStore{
	return &PostgresTaxStore{
		db: db,
	}
}

func (pg *PostgresTaxStore) CreateTax(tax *Tax) (*Tax,error) {
	tx, err := pg.db.Begin()
	if err != nil {
		return nil, err
	}

	const query = `INSERT INTO taxes(name, label, tax_percentage) VALUES ($1, $2, $3) RETURNING id`

	err = tx.QueryRow(query, &tax.Name, &tax.Label, &tax.TaxPercentage).Scan(&tax.Id)

	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return tax, nil
}

func (pg *PostgresTaxStore) GetTaxByID(id  int64) (*Tax, error) {
	var tax = Tax{}
	const query = `SELECT id, name, label, tax_percentage FROM taxes WHERE id = $1`;
	err := pg.db.QueryRow(query,id).Scan(&tax.Id, &tax.Name, &tax.Label, &tax.TaxPercentage);

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &tax, nil
}

func (pg *PostgresTaxStore) GetTaxes() (*[]Tax, error) {
	const query = `SELECT id, name, label, tax_percentage FROM taxes LIMIT 100`

	rows, err := pg.db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var taxes []Tax

	for rows.Next() {
		var tax = Tax{}
		err = rows.Scan(&tax.Id, &tax.Name, &tax.Label, &tax.TaxPercentage)

		if err != nil {
			return nil, err
		}

		taxes = append(taxes,tax)
	}

	return &taxes, nil
}
