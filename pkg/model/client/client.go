package client

import (
	"context"
	"time"

	"github.com/lpiegas25/go_store/internal/data"
)

type Client struct {
	ID           uint      `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Lastname     string    `json:"lastname,omitempty"`
	PrimaryPhone string    `json:"primary_phone,omitempty"`
	SecondPhone  string    `json:"second_phone,omitempty"`
	Address      string    `json:"address,omitempty"`
	Email        string    `json:"email,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	AccountId    uint      `json:"account_id,omitempty"`
}

type Repository struct {
	Data *data.Data
}

func (cr Repository) GetAll(ctx context.Context) ([]Client, error) {
	q := `SELECT id, account_id, name, lastname, primary_phone, second_phone, address, email, created_at, updated_at
        FROM clients`
	rows, err := cr.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clients []Client
	for rows.Next() {
		var c Client
		rows.Scan(&c.ID, &c.AccountId, &c.Name, &c.Lastname, &c.PrimaryPhone, &c.SecondPhone, &c.Address, &c.Email, &c.CreatedAt, &c.UpdatedAt)
		clients = append(clients, c)
	}
	return clients, nil
}

func (cr Repository) GetOne(ctx context.Context, id uint) (Client, error) {

	q := `SELECT id, account_id, name, lastname, primary_phone, second_phone, address, email, created_at, updated_at
        FROM clients
		WHERE id=$1; 
		`
	row := cr.Data.DB.QueryRowContext(ctx, q, id)

	var c Client
	err := row.Scan(&c.ID, &c.AccountId, &c.Name, &c.Lastname, &c.PrimaryPhone, &c.SecondPhone, &c.Address, &c.Email, &c.CreatedAt, &c.UpdatedAt)
	if err != nil {
		return Client{}, err
	}
	return c, nil
}

func (cr Repository) Create(ctx context.Context, c *Client) error {
	q := `INSERT INTO clients (account_id, name, lastname, primary_phone, second_phone, address, email, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id`

	stmt, err := cr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, c.AccountId, c.Name, c.Lastname, c.PrimaryPhone, c.SecondPhone, c.Address, c.Email, time.Now(), time.Now())

	err = row.Scan(&c.ID)
	if err != nil {
		return err
	}
	return nil
}

func (cr Repository) Update(ctx context.Context, id uint, c Client) error {
	q := `UPDATE clients set account_id=$1, name=$2, lastname=$3, primary_phone=$4, second_phone=$5, address=$6, email=$7, updated_at=$8
			WHERE id=$9`
	stmt, err := cr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, c.AccountId, c.Name, c.Lastname, c.PrimaryPhone, c.SecondPhone, c.Address, c.Email, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}

func (cr Repository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM clients WHERE id=$1;`

	stmt, err := cr.Data.DB.PrepareContext(ctx, q)
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
