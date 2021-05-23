package client

import (
	"context"
	"github.com/lpiegas25/go_store/internal/data"
	"time"
)

type Client struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	Lastname     string    `json:"lastname"`
	PrimaryPhone string    `json:"primary_phone"`
	SecondPhone  string    `json:"second_phone"`
	Address      string    `json:"address"`
	Email        string    `json:"email"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`
	AccountId    uint      `json:"account_id"`
}

type Repository interface {
	GetAll(ctx context.Context) ([]Client, error)
	GetOne(ctx context.Context, id uint) (Client, error)
	Create(ctx context.Context, client *Client) error
	Update(ctx context.Context, id uint, client Client) error
	Delete(ctx context.Context, id uint) error
}

type ClientRepository struct {
	Data *data.Data
}

func (cr ClientRepository) GetAll(ctx context.Context) ([]Client, error) {
	//ToDo: Implement get all client
	panic("implement me")
}

func (cr ClientRepository) GetOne(ctx context.Context, id uint) (Client, error) {

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

func (cr ClientRepository) Create(ctx context.Context, c *Client) error {
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

func (cr ClientRepository) Update(ctx context.Context, id uint, client Client) error {
	//ToDo: Implement update client
	panic("implement me")
}

func (cr ClientRepository) Delete(ctx context.Context, id uint) error {
	//ToDo: Implement delete client
	panic("implement me")
}
