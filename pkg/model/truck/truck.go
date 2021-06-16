package truck

import (
	"context"
	"time"

	"github.com/lpiegas25/go_store/internal/data"
)

type Truck struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Repository struct {
	Data *data.Data
}

func (wr Repository) GetAll(ctx context.Context) ([]Truck, error) {
	q := `SELECT id, name, created_at, updated_at
        FROM trucks`
	rows, err := wr.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trucks []Truck
	for rows.Next() {
		var truck Truck
		rows.Scan(&truck.ID, &truck.Name, &truck.CreatedAt, &truck.UpdatedAt)
		trucks = append(trucks, truck)
	}
	return trucks, nil
}

func (wr Repository) GetOne(ctx context.Context, id uint) (Truck, error) {
	q := `SELECT id, name, created_at, updated_at
        FROM trucks
		WHERE id=$1; 
		`
	row := wr.Data.DB.QueryRowContext(ctx, q, id)

	var truck Truck
	err := row.Scan(&truck.ID, &truck.Name, &truck.CreatedAt, &truck.UpdatedAt)
	if err != nil {
		return Truck{}, err
	}
	return truck, nil
}

func (wr Repository) Create(ctx context.Context, truck *Truck) error {
	q := `INSERT INTO trucks (name, created_at, updated_at)
		VALUES ($1, $2, $3)
		RETURNING id`

	stmt, err := wr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, truck.Name, time.Now(), time.Now())

	err = row.Scan(&truck.ID)
	if err != nil {
		return err
	}
	return nil
}

func (wr Repository) Update(ctx context.Context, id uint, truck Truck) error {
	q := `UPDATE trucks set name=$1, updated_at=$2
			WHERE id=$3`
	stmt, err := wr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, truck.Name, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}

func (wr Repository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM trucks WHERE id=$1;`

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
