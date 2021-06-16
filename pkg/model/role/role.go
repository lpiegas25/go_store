package role

import (
	"context"
	"time"

	"github.com/lpiegas25/go_store/internal/data"
)

type Role struct {
	ID          uint      `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

type Repository struct {
	Data *data.Data
}

func (rr Repository) GetAll(ctx context.Context) ([]Role, error) {
	q := `SELECT id, name, description, created_at, updated_at
        FROM roles`
	rows, err := rr.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []Role
	for rows.Next() {
		var role Role
		rows.Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
		roles = append(roles, role)
	}
	return roles, nil
}

func (rr Repository) GetOne(ctx context.Context, id uint) (Role, error) {
	q := `SELECT id, name, description, created_at, updated_at
        FROM roles
		WHERE id=$1; 
		`
	row := rr.Data.DB.QueryRowContext(ctx, q, id)

	var role Role
	err := row.Scan(&role.ID, &role.Name, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		return Role{}, err
	}
	return role, nil
}

func (rr Repository) Create(ctx context.Context, role *Role) error {
	q := `INSERT INTO roles (name, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	stmt, err := rr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, role.Name, role.Description, time.Now(), time.Now())

	err = row.Scan(&role.ID)
	if err != nil {
		return err
	}
	return nil
}

func (rr Repository) Update(ctx context.Context, id uint, role Role) error {
	q := `UPDATE roles set name=$1, description=$2, updated_at=$3
			WHERE id=$4`
	stmt, err := rr.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, role.Name, role.Description, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}

func (rr Repository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM roles WHERE id=$1;`

	stmt, err := rr.Data.DB.PrepareContext(ctx, q)
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
