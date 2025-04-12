package postgre

import (
	"context"
	"database/sql"
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

func (m *MagazineModel) Get(name, city string) (int, error) {
	return 0, nil
}
