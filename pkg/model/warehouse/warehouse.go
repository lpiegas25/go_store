package warehouse

import (
	"context"
	"github.com/lpiegas25/go_store/internal/data"
	"time"
)

type Warehouse struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type WarehouseRepository struct {
	Data *data.Data
}

func (wr WarehouseRepository) GetAll(ctx context.Context) ([]Warehouse, error) {
	q := `SELECT id, name, created_at, updated_at
        FROM warehouses`
	rows, err := wr.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var warehouses []Warehouse
	for rows.Next() {
		var warehouse Warehouse
		rows.Scan(&warehouse.ID, &warehouse.Name, &warehouse.CreatedAt, &warehouse.UpdatedAt)
		warehouses = append(warehouses, warehouse)
	}
	return warehouses, nil
}

func (wr WarehouseRepository) GetOne(ctx context.Context, id uint) (Warehouse, error) {
	q := `SELECT id, name, created_at, updated_at
        FROM warehouses
		WHERE id=$1; 
		`
	row := wr.Data.DB.QueryRowContext(ctx, q, id)

	var warehouse Warehouse
	err := row.Scan(&warehouse.ID, &warehouse.Name, &warehouse.CreatedAt, &warehouse.UpdatedAt)
	if err != nil {
		return Warehouse{}, err
	}
	return warehouse, nil
}

func (wr WarehouseRepository) Create(ctx context.Context, warehouse *Warehouse) error {
	q := `INSERT INTO warehouses (name, created_at, updated_at)
		VALUES ($1, $2, $3)
		RETURNING id`

	stmt, err := wr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, warehouse.Name, time.Now(), time.Now())

	err = row.Scan(&warehouse.ID)
	if err != nil {
		return err
	}
	return nil
}

func (wr WarehouseRepository) Update(ctx context.Context, id uint, warehouse Warehouse) error {
	q := `UPDATE warehouses set name=$1, updated_at=$2
			WHERE id=$3`
	stmt, err := wr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, warehouse.Name, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}

func (wr WarehouseRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM warehouses WHERE id=$1;`

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
