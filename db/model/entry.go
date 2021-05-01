package model

type CreateEntryPayload struct {
	AccountID uint64 `json:"account_id"`
	Amount    int64  `json:"amount"`
}

const createEntry = `INSERT INTO entries (account_id, amount) 
VALUES ($1, $2) 
RETURNING id, account_id, amount, created_at`

func (r *Repository) CreateEntry(arg CreateEntryPayload) (Entry, error) {
	row := r.db.QueryRow(createEntry, arg.AccountID, arg.Amount)
	var entry Entry
	err := row.Scan(
		&entry.ID,
		&entry.AccountID,
		&entry.Amount,
		&entry.CreatedAt,
	)
	return entry, err
}

const getEntry = `SELECT id, account_id, amount, created_at 
FROM entries 
WHERE id = $1 
LIMIT 1`

func (r *Repository) GetEntry(id uint64) (Entry, error) {
	row := r.db.QueryRow(getEntry, id)

	var entry Entry
	err := row.Scan(
		&entry.ID,
		&entry.AccountID,
		&entry.Amount,
		&entry.CreatedAt,
	)

	return entry, err
}

type ListEntryParams struct {
	AccountID uint64 `json:"account_id"`
	Limit     int32  `json:"limit"`
	Offset    int32  `json:"offset"`
}

const listEntry = `SELECT id, account_id, amount, created_at 
FROM entries 
WHERE account_id = $1 
LIMIT $2 OFFSET $3`

func (r *Repository) ListEntry(args ListEntryParams) ([]Entry, error) {
	rows, err := r.db.Query(listEntry, args.AccountID, args.Limit, args.Offset)
	if err != nil {
		return nil, err
	}
	var items []Entry
	for rows.Next() {
		var i Entry
		err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.Amount,
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
