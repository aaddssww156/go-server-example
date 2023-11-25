package services

import (
	"context"
	"fmt"
	"time"
)

type Coffe struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Roast     string    `json:"roast,omitempty"`
	Image     string    `json:"image,omitempty"`
	Region    string    `json:"region,omitempty"`
	Price     float32   `json:"price,omitempty"`
	GrindUnit int16     `json:"grind_unit,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (c *Coffe) GetAllCoffes() ([]*Coffe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `SELECT id, name, image, roast, region, price, grind_unit, created_at, updated_at FROM coffees`
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error getting all coffees: %s", err.Error())
	}

	var coffees []*Coffe
	for rows.Next() {
		var coffee Coffe

		err := rows.Scan(
			&coffee.ID,
			&coffee.Name,
			&coffee.Image,
			&coffee.Roast,
			&coffee.Region,
			&coffee.Price,
			&coffee.GrindUnit,
			&coffee.CreatedAt,
			&coffee.UpdatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("error scanning coffee row: %s", err.Error())
		}

		coffees = append(coffees, &coffee)
	}

	return coffees, nil
}

func (c *Coffe) CreateCoffee(coffee Coffe) (*Coffe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `INSERT INTO coffees (name, image, region, roast, price, grind_unit, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) returning *`

	_, err := db.ExecContext(ctx, query, coffee.Name, coffee.Image, coffee.Region, coffee.Roast, coffee.Price, coffee.GrindUnit, time.Now(), time.Now())
	if err != nil {
		return nil, fmt.Errorf("error inserting coffe: %s", err.Error())
	}

	return &coffee, nil
}
