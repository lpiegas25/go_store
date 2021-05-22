package account

import (
	"context"
	"github.com/lpiegas25/go_store/internal/data"
	"time"
)

type Account struct {
	ID uint `json:"id"`
	ActualAmount float64 `json:"actual_amount"`
	PreviousAmount float64 `json:"previous_amount"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}


type Repository interface {
	GetAll(ctx context.Context) ([]Account, error)
	GetOne(ctx context.Context, id uint) (Account, error)
	Create(ctx context.Context, account *Account) error
	Update(ctx context.Context, id uint, account Account) error
	Delete(ctx context.Context, id uint) error
}

type AccountRepository struct {
	Data *data.Data
}

func (a AccountRepository) GetAll(ctx context.Context) ([]Account, error) {
	panic("implement me")
}

func (a AccountRepository) GetOne(ctx context.Context, id uint) (Account, error) {
	panic("implement me")
}

func (a AccountRepository) Create(ctx context.Context, ac *Account) error {
	q := `INSERT INTO accounts (actual_amount, previous_amount, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id`
	
	stmt, err := a.Data.DB.PrepareContext(ctx,q)
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

func (a AccountRepository) Update(ctx context.Context, id uint, account Account) error {
	panic("implement me")
}

func (a AccountRepository) Delete(ctx context.Context, id uint) error {
	panic("implement me")
}
