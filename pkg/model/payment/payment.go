package payment

import (
	"context"
	"time"

	"github.com/lpiegas25/go_store/internal/data"
)

type Payment struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Repository struct {
	Data *data.Data
}

func (wr Repository) GetAll(ctx context.Context) ([]Payment, error) {
	q := `SELECT id, name, created_at, updated_at
        FROM payments`
	rows, err := wr.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []Payment
	for rows.Next() {
		var payment Payment
		rows.Scan(&payment.ID, &payment.Name, &payment.CreatedAt, &payment.UpdatedAt)
		payments = append(payments, payment)
	}
	return payments, nil
}

func (wr Repository) GetOne(ctx context.Context, id uint) (Payment, error) {
	q := `SELECT id, name, created_at, updated_at
        FROM payments
		WHERE id=$1; 
		`
	row := wr.Data.DB.QueryRowContext(ctx, q, id)

	var payment Payment
	err := row.Scan(&payment.ID, &payment.Name, &payment.CreatedAt, &payment.UpdatedAt)
	if err != nil {
		return Payment{}, err
	}
	return payment, nil
}

func (wr Repository) Create(ctx context.Context, payment *Payment) error {
	q := `INSERT INTO payments (name, created_at, updated_at)
		VALUES ($1, $2, $3)
		RETURNING id`

	stmt, err := wr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, payment.Name, time.Now(), time.Now())

	err = row.Scan(&payment.ID)
	if err != nil {
		return err
	}
	return nil
}

func (wr Repository) Update(ctx context.Context, id uint, payment Payment) error {
	q := `UPDATE payments set name=$1, updated_at=$2
			WHERE id=$3`
	stmt, err := wr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, payment.Name, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}

func (wr Repository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM payments WHERE id=$1;`

	stmt, err := wr.Data.DB.PrepareContext(ctx, q)
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
