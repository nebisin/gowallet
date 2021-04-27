package model

type CreateAccountPayload struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

const createAccount = `INSERT INTO accounts (
owner, balance, currency
) VALUES (
$1, $2, $3
) RETURNING id, owner, balance, currency, created_at`

func (r Repository) CreateAccount(payload CreateAccountPayload) (Account, error) {
	row := r.db.QueryRow(createAccount, payload.Owner, payload.Balance, payload.Currency)

	var account Account
	err := row.Scan(
		&account.ID,
		&account.Owner,
		&account.Balance,
		&account.Currency,
		&account.CreatedAt,
	)

	return account, err
}

const getAccount = `SELECT id, owner, balance, currency, created_at 
FROM accounts 
WHERE id = $1 
LIMIT 1`

func (r Repository) GetAccount(id uint64) (Account, error) {
	row := r.db.QueryRow(getAccount, id)

	var account Account
	err := row.Scan(
		&account.ID,
		&account.Owner,
		&account.Balance,
		&account.Currency,
		&account.CreatedAt,
	)

	return account, err
}

type ListAccountParams struct {
	Owner  string `json:"owner"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

const listAccount = `SELECT id, owner, balance, currency, created_at
FROM accounts
WHERE owner = $1 
LIMIT $2 OFFSET $3`

func (r Repository) ListAccount(arg ListAccountParams) ([]Account, error) {
	rows, err := r.db.Query(listAccount, arg.Owner, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}

	var items []Account
	for rows.Next() {
		var i Account
		err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Balance,
			&i.Currency,
			&i.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		items = append(items, i)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

type UpdateAccountParams struct {
	ID      uint64 `json:"id"`
	Balance int64  `json:"balance"`
}

const updateAccount = `UPDATE accounts 
SET balance = $2 
WHERE id = $1 
RETURNING id, owner, balance, currency, created_at`

func (r Repository) UpdateAccount(arg UpdateAccountParams) (Account, error) {
	row := r.db.QueryRow(updateAccount, arg.ID, arg.Balance)

	var account Account
	err := row.Scan(
		&account.ID,
		&account.Owner,
		&account.Balance,
		&account.Currency,
		&account.CreatedAt,
	)
	return account, err

}

const deleteAccount = `DELETE FROM accounts WHERE id = $1`

func (r Repository) DeleteAccount(id uint64) error {
	_, err := r.db.Exec(deleteAccount, id)

	return err
}
