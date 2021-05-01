package model

type CreateTransferPayload struct {
	FromAccountID uint64 `json:"from_account_id"`
	ToAccountID   uint64 `json:"to_account_id"`
	Amount        int64  `json:"amount"`
}

const createTransfer = `INSERT INTO transfers (from_account_id, to_account_id, amount) 
VALUES ($1, $2, $3)
RETURNING id, from_account_id, to_account_id, amount, created_at`

func (r *Repository) CreateTransfer(args CreateTransferPayload) (Transfer, error) {
	row := r.db.QueryRow(createTransfer, args.FromAccountID, args.ToAccountID, args.Amount)

	var transfer Transfer
	err := row.Scan(
		&transfer.ID,
		&transfer.FromAccountID,
		&transfer.ToAccountID,
		&transfer.Amount,
		&transfer.CreatedAt,
	)

	return transfer, err
}

const getTransfer = `SELECT id, from_account_id, to_account_id, amount, created_at 
FROM transfers WHERE id = $1 LIMIT 1`

func (r *Repository) GetTransfer(id uint64) (Transfer, error) {
	row := r.db.QueryRow(getTransfer, id)
	var transfer Transfer
	err := row.Scan(
		&transfer.ID,
		&transfer.FromAccountID,
		&transfer.ToAccountID,
		&transfer.Amount,
		&transfer.CreatedAt,
	)
	return transfer, err
}

const listTransfer = `SELECT id, from_account_id, to_account_id, amount, created_at 
FROM transfers WHERE from_account_id = $1 OR to_account_id = $2 
LIMIT $3 OFFSET $4`

type ListTransferParams struct {
	FromAccountID uint64 `json:"from_account_id"`
	ToAccountID   uint64 `json:"to_account_id"`
	Limit         int32  `json:"limit"`
	Offset        int32  `json:"offset"`
}

func (r *Repository) ListTransfer(arg ListTransferParams) ([]Transfer, error) {
	rows, err := r.db.Query(listTransfer, arg.FromAccountID, arg.ToAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}

	var transfers []Transfer
	for rows.Next() {
		var transfer Transfer
		err := rows.Scan(
			&transfer.ID,
			&transfer.FromAccountID,
			&transfer.ToAccountID,
			&transfer.Amount,
			&transfer.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		transfers = append(transfers, transfer)

	}
	return transfers, nil
}
