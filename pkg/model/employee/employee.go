package employee

import (
	"context"
	"github.com/lpiegas25/go_store/internal/data"
	"time"
)

type Employee struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Lastname  string    `json:"lastname,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	RoleId    uint      `json:"role_id,omitempty"`
}

type Repository interface {
	GetAll(ctx context.Context) ([]Employee, error)
	GetOne(ctx context.Context, id uint) (Employee, error)
	Create(ctx context.Context, employee *Employee) error
	Update(ctx context.Context, id uint, employee Employee) error
	Delete(ctx context.Context, id uint) error
}

type EmployeeRepository struct {
	Data *data.Data
}

func (er EmployeeRepository) GetAll(ctx context.Context) ([]Employee, error) {
	q := `SELECT id, role_id, name, lastname, phone, created_at, updated_at
        FROM employees`
	rows, err := er.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []Employee
	for rows.Next() {
		var e Employee
		rows.Scan(&e.ID, &e.RoleId, &e.Name, &e.Lastname, &e.Phone, &e.CreatedAt, &e.UpdatedAt)
		employees = append(employees, e)
	}
	return employees, nil
}

func (er EmployeeRepository) GetOne(ctx context.Context, id uint) (Employee, error) {

	q := `SELECT id, role_id, name, lastname, phone, created_at, updated_at
        FROM employees
		WHERE id=$1; 
		`
	row := er.Data.DB.QueryRowContext(ctx, q, id)

	var e Employee
	err := row.Scan(&e.ID, &e.RoleId, &e.Name, &e.Lastname, &e.Phone, &e.CreatedAt, &e.UpdatedAt)
	if err != nil {
		return Employee{}, err
	}
	return e, nil
}

func (er EmployeeRepository) Create(ctx context.Context, e *Employee) error {
	q := `INSERT INTO employees (role_id, name, lastname, phone, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	stmt, err := er.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, e.RoleId, e.Name, e.Lastname, e.Phone, time.Now(), time.Now())

	err = row.Scan(&e.ID)
	if err != nil {
		return err
	}
	return nil
}

func (er EmployeeRepository) Update(ctx context.Context, id uint, e Employee) error {
	q := `UPDATE employees set role_id=$1, name=$2, lastname=$3, phone=$4, updated_at=$5
			WHERE id=$6`
	stmt, err := er.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, e.RoleId, e.Name, e.Lastname, e.Phone, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}

func (er EmployeeRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM employees WHERE id=$1;`

	stmt, err := er.Data.DB.PrepareContext(ctx, q)
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
