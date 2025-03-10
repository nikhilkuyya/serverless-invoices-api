package store

import "database/sql"

type BankAccount struct {
	ID int `json:"id"`
	AccountNickName string `json:"account_nick_name"`
	BankName string `json:"bank_name"`
	BankAccountNumber string `json:"bank_account_number"`
	BankIfscCode string `json:"bank_ifsc_code"`
	BankDescription string `json:"bank_description"`
}

type BankAccountStore interface {
	CreateBankAccount(bankAccount *BankAccount) (*BankAccount,error)
	GetBankAccountByID(id int64) (*BankAccount,error)
	GetAllBankAccounts() (*[]BankAccount,error)
}

type PostgresBankAccountStore struct {
	db *sql.DB
}

func NewPostgresBankAccountStore(db *sql.DB) *PostgresBankAccountStore {
	return &PostgresBankAccountStore{db: db}
}

func (pg *PostgresBankAccountStore) CreateBankAccount(bankAccount *BankAccount) (*BankAccount, error) {
	tx, err := pg.db.Begin()
	if err != nil {
		return nil, err
	}
	// NOT SURE, completely understood this defer working
	defer tx.Rollback()

	query := `
		INSERT INTO bank_accounts
		(account_nick_name, bank_name, bank_account_number, bank_ifsc_code, bank_description )
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`

	err = tx.QueryRow(query, bankAccount.AccountNickName,bankAccount.BankName, bankAccount.BankAccountNumber, bankAccount.BankIfscCode, bankAccount.BankDescription).Scan(&bankAccount.ID)

	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return bankAccount,nil
}

func (pg *PostgresBankAccountStore) GetBankAccountByID(id int64) (*BankAccount, error) {
	bankAccount:= BankAccount{};
	tx, err := pg.db.Begin()
	if err != nil {
		return nil, err
	}

	query := "SELECT id,account_nick_name, bank_name, bank_account_number, bank_ifsc_code,bank_description from bank_accounts WHERE id = $1"

	err = tx.QueryRow(query, id).Scan(&bankAccount.ID, &bankAccount.AccountNickName, &bankAccount.BankName, &bankAccount.BankAccountNumber,&bankAccount.BankIfscCode,&bankAccount.BankDescription);

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &bankAccount, nil
}


func (pg *PostgresBankAccountStore) GetAllBankAccounts() (*[]BankAccount,error) {
	query := "SELECT id,account_nick_name, bank_name, bank_account_number, bank_ifsc_code, bank_description from bank_accounts LIMIT 10"

	var bankAccounts []BankAccount;

	rows, err := pg.db.Query(query);

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var bankAccount BankAccount
		err = rows.Scan(&bankAccount.ID, &bankAccount.AccountNickName, &bankAccount.BankName,
		&bankAccount.BankAccountNumber, &bankAccount.BankIfscCode, &bankAccount.BankDescription)
		if err != nil {
			return nil, err
		}
		bankAccounts = append(bankAccounts, bankAccount);
	}

	return &bankAccounts,nil
}
