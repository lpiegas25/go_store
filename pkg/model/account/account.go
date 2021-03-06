package account

import (
	"context"
	"time"

	"github.com/lpiegas25/go_store/internal/data"
)

type Account struct {
	ID             uint      `json:"id,omitempty"`
	ActualAmount   float64   `json:"actual_amount,omitempty"`
	PreviousAmount float64   `json:"previous_amount,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}

type Repository struct {
	Data *data.Data
}

func (a Repository) GetOne(ctx context.Context, id uint) (Account, error) {
	q := `SELECT id, actual_amount, previous_amount, created_at, updated_at
        FROM accounts
		WHERE id=$1; 
		`
	row := a.Data.DB.QueryRowContext(ctx, q, id)

	var ac Account
	err := row.Scan(&ac.ID, &ac.ActualAmount, &ac.PreviousAmount, &ac.CreatedAt, &ac.UpdatedAt)
	if err != nil {
		return Account{}, err
	}
	return ac, nil
}

func (a Repository) Create(ctx context.Context, ac *Account) error {
	q := `INSERT INTO accounts (actual_amount, previous_amount, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	stmt, err := a.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, ac.ActualAmount, ac.PreviousAmount, time.Now(), time.Now())

	err = row.Scan(&ac.ID)
	if err != nil {
		return err
	}
	return nil
}

func (a Repository) Update(ctx context.Context, id uint, ac Account) error {
	q := `UPDATE accounts set actual_amount=$1, previous_amount=actual_amount, updated_at=$2
			WHERE id=$3`
	stmt, err := a.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, ac.ActualAmount, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}

func (a Repository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM accounts WHERE id=$1;`

	stmt, err := a.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
