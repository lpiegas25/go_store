package delivery

import (
	"context"
	"fmt"
	"github.com/lpiegas25/go_store/internal/data"
	"github.com/lpiegas25/go_store/pkg/model/employee"
	"time"
)

// Delivery
// When a delivery is created,
// comes with an array of employees that are the ones that are on board of the truck
type Delivery struct {
	ID          uint      `json:"id,omitempty"`
	TruckId     uint      `json:"truck_id,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	EmployeesId []uint    `json:"employees_id,omitempty"`
}

type DeliveryDTO struct {
	ID        uint                `json:"id,omitempty"`
	TruckId   uint                `json:"truck_id,omitempty"`
	CreatedAt time.Time           `json:"created_at,omitempty"`
	UpdatedAt time.Time           `json:"updated_at,omitempty"`
	Employees []employee.Employee `json:"employees,omitempty"`
	TruckName string              `json:"truck_name,omitempty"`
}

type Repository struct {
	Data *data.Data
}

func (r Repository) Create(ctx context.Context, delivery *Delivery) error {

	tx, err := r.Data.DB.BeginTx(ctx, nil)

	q := `INSERT INTO deliveries (truck_id, created_at, updated_at)
		VALUES ($1, $2, $3)
		RETURNING id
		`
	row := tx.QueryRowContext(ctx, q, delivery.TruckId, time.Now(), time.Now())

	err = row.Scan(&delivery.ID)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf(`error on delivery.Create insert --> %q`, err)
	}
	if len(delivery.EmployeesId) > 0 {
		q2 := `INSERT INTO deliveries_employees (delivery_id, employee_id, created_at, updated_at)
				VALUES ($1, $2, $3, $4)`

		for _, id := range delivery.EmployeesId {
			_, err = tx.ExecContext(ctx, q2, delivery.ID, id, time.Now(), time.Now())
		}
		if err != nil {
			_ = tx.Rollback()
			return fmt.Errorf(`error on delivery.Create insert into deliveries_employees --> %q`, err)
		}
	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) GetOne(ctx context.Context, id uint) (DeliveryDTO, error) {
	queryGetDelivery := `SELECT d.id, truck_id, d.created_at, d.updated_at, t.name
        FROM deliveries d
		INNER JOIN trucks t ON d.truck_id = t.id
		WHERE d.id=$1 
		`

	row := r.Data.DB.QueryRowContext(ctx, queryGetDelivery, id)

	var deliveryDTO DeliveryDTO
	err := row.Scan(&deliveryDTO.ID, &deliveryDTO.TruckId, &deliveryDTO.CreatedAt, &deliveryDTO.UpdatedAt, &deliveryDTO.TruckName)
	if err != nil {
		return DeliveryDTO{}, err
	}

	queryGetDeliveryEmployees := `
		SELECT e.id, e.role_id, e.name, e.lastname, e.phone, e.created_at, e.updated_at 
		FROM employees e
		INNER JOIN deliveries_employees de on e.id = de.employee_id
		WHERE de.delivery_id = $1`

	rows, err := r.Data.DB.QueryContext(ctx, queryGetDeliveryEmployees, deliveryDTO.ID)
	if err != nil {
		return DeliveryDTO{}, err
	}
	var employees []employee.Employee
	for rows.Next() {
		var e employee.Employee
		err = rows.Scan(&e.ID, &e.RoleId, &e.Name, &e.Lastname, &e.Phone, &e.CreatedAt, &e.UpdatedAt)
		employees = append(employees, e)
	}
	deliveryDTO.Employees = employees

	return deliveryDTO, nil
}
