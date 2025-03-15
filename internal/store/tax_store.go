package store

import "database/sql"

type Tax struct {
	Id int64 `json:"id"`
	Label string `json:"label"`
	TaxPercentage string `json:"tax_percentage"`
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
	return nil, nil
}

func (pg *PostgresTaxStore) GetTaxByID(id  int64) (*Tax, error) {
	return nil, nil
}

func (pg *PostgresTaxStore) GetTaxes() (*[]Tax, error) {
	return nil, nil
}
