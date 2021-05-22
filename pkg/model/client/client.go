package client

import (
	"context"
	"github.com/lpiegas25/go_store/internal/data"
	"time"
)

type Client struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Lastname string `json:"lastname"`
	PrimaryPhone string `json:"primary_phone"`
	SecondPhone string `json:"second_phone"`
	Address string `json:"address"`
	Email string `json:"email"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	AccountId uint `json:"account_id"`
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

func (c ClientRepository) GetAll(ctx context.Context) ([]Client, error) {
	panic("implement me")
}

func (c ClientRepository) GetOne(ctx context.Context, id uint) (Client, error) {
	panic("implement me")
}

func (c ClientRepository) Create(ctx context.Context, client *Client) error {
	panic("implement me")
}

func (c ClientRepository) Update(ctx context.Context, id uint, client Client) error {
	panic("implement me")
}

func (c ClientRepository) Delete(ctx context.Context, id uint) error {
	panic("implement me")
}

