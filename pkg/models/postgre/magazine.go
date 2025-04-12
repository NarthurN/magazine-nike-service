package postgre

import (
	"context"
	"database/sql"
	"microservice1/pkg/models"
)

type MagazineModel struct {
	DB *sql.DB
}

func (m *MagazineModel) Insert(name, city string) (int, error) {
	var id int
	stmt := `INSERT INTO magazines (name, city)
			VALUES ($1, $2) RETURNING id`
	row := m.DB.QueryRowContext(context.Background(), stmt, name, city)
	return id, row.Scan(&id)
}

func (m *MagazineModel) Get(city string) ([]models.Magazine, error) {
	stmt := `SELECT id, name, city FROM magazines WHERE city=$1`
	rows, err := m.DB.QueryContext(context.Background(), stmt, city)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []models.Magazine

	for rows.Next() {
		magazine := models.Magazine{}
		err := rows.Scan(&magazine.Id, &magazine.Name, &magazine.City)
		if err != nil {
			return nil, err
		}
		res = append(res, magazine)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
