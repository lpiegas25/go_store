package invoice

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/lpiegas25/go_store/internal/data"
	"time"
)

type Line struct {
	ID        uint      `json:"id,omitempty"`
	ItemID    uint      `json:"item_id,omitempty"`
	InvoiceID uint      `json:"invoice_id,omitempty"`
	Quantity  float64   `json:"quantity,omitempty"`
	Amount    float64   `json:"amount,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Invoice struct {
	ID             uint      `json:"id,omitempty"`
	PaymentID      uint      `json:"payment_id,omitempty"`
	ClientID       uint      `json:"client_id,omitempty"`
	TotalAmount    float64   `json:"total_amount,omitempty"`
	DiscountAmount float64   `json:"discount_amount,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	Lines          []Line    `json:"lines,omitempty"`
}

type Repository struct {
	Data *data.Data
}

func (r Repository) GetAll(ctx context.Context) ([]Invoice, error) {
	q := `SELECT id, payment_id, client_id, total_amount, discount_amount, created_at, updated_at
        FROM invoices`
	rows, err := r.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []Invoice
	for rows.Next() {
		var invoice Invoice
		err = rows.Scan(&invoice.ID, &invoice.PaymentID, &invoice.ClientID, &invoice.TotalAmount, &invoice.DiscountAmount, &invoice.CreatedAt, &invoice.UpdatedAt)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, invoice)
	}
	return invoices, nil
}

func (r Repository) GetOne(ctx context.Context, id uint) (Invoice, error) {
	q := `SELECT id, payment_id, client_id, total_amount, discount_amount, created_at, updated_at
        FROM invoices
		WHERE id=$1 
		`

	row := r.Data.DB.QueryRowContext(ctx, q, id)

	var invoice Invoice
	err := row.Scan(&invoice.ID, &invoice.PaymentID, &invoice.ClientID, &invoice.TotalAmount, &invoice.DiscountAmount, &invoice.CreatedAt, &invoice.UpdatedAt)
	if err != nil {
		return Invoice{}, err
	}
	invoice.Lines, err = r.GetManyLineByInvoiceId(ctx, invoice.ID)
	if err != nil {
		return Invoice{}, err
	}

	return invoice, nil
}

func (r Repository) Create(ctx context.Context, invoice *Invoice) error {

	tx, err := r.Data.DB.BeginTx(ctx, nil)

	q := `INSERT INTO invoices (payment_id, client_id, total_amount, discount_amount, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
		`
	row := tx.QueryRowContext(ctx, q, invoice.PaymentID, invoice.ClientID, invoice.TotalAmount, invoice.DiscountAmount, time.Now(), time.Now())

	err = row.Scan(&invoice.ID)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf(`error on invoice.Create insert --> %q`, err)
	}
	if len(invoice.Lines) > 0 {
		for _, line := range invoice.Lines {
			line.InvoiceID = invoice.ID
			err = r.CreateLine(ctx, tx, &line)
		}
		if err != nil {
			_ = tx.Rollback()
			return fmt.Errorf(`error on invoice.CreateLine --> insert --> %q`, err)
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil

}

func (r Repository) Update(ctx context.Context, id uint, invoice Invoice) error {
	q := `UPDATE invoices set payment_id=$1, client_id=$2, total_amount=$3, discount_amount=$4, updated_at=$5
			WHERE id=$6`
	stmt, err := r.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, invoice.PaymentID, invoice.ClientID, invoice.TotalAmount, invoice.DiscountAmount, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM invoices WHERE id=$1;`

	stmt, err := r.Data.DB.PrepareContext(ctx, q)
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

func (r Repository) GetOneLine(ctx context.Context, id uint) (Line, error) {
	q := `SELECT id, item_id, invoice_id, quantity, amount, created_at, updated_at
        FROM invoice_lines
		WHERE id=$1 
		`
	row := r.Data.DB.QueryRowContext(ctx, q, id)

	var line Line
	err := row.Scan(&line.ID, &line.ItemID, &line.InvoiceID, &line.Quantity, &line.Amount, &line.CreatedAt, &line.UpdatedAt)
	if err != nil {
		return Line{}, err
	}
	return line, nil
}

func (r Repository) GetManyLineByInvoiceId(ctx context.Context, invoiceId uint) ([]Line, error) {
	q := `SELECT id, item_id, invoice_id, quantity, amount, created_at, updated_at
        FROM invoice_lines
        WHERE invoice_id = $1`
	rows, err := r.Data.DB.QueryContext(ctx, q, invoiceId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lines []Line
	for rows.Next() {
		var line Line
		err = rows.Scan(&line.ID, &line.ItemID, &line.InvoiceID, &line.Quantity, &line.Amount, &line.CreatedAt, &line.UpdatedAt)
		if err != nil {
			return nil, err
		}
		lines = append(lines, line)
	}
	return lines, nil
}

func (r Repository) CreateLine(ctx context.Context, tx *sql.Tx, line *Line) error {
	q := `INSERT INTO invoice_lines (item_id, invoice_id, quantity, amount, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
		`

	row := tx.QueryRowContext(ctx, q, line.ItemID, line.InvoiceID, line.Quantity, line.Amount, time.Now(), time.Now())
	err := row.Scan(&line.ID)

	if err != nil {
		return err
	}
	return nil
}

func (r Repository) UpdateLine(ctx context.Context, id uint, line Line) error {
	q := `UPDATE invoice_lines set item_id=$1, invoice_id=$2, quantity=$3, amount=$4, updated_at=$5
			WHERE id=$6`
	stmt, err := r.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, line.ItemID, line.InvoiceID, line.Quantity, line.Amount, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}

func (r Repository) DeleteLine(ctx context.Context, id uint) error {
	q := `DELETE FROM invoice_lines WHERE id=$1;`

	stmt, err := r.Data.DB.PrepareContext(ctx, q)
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
