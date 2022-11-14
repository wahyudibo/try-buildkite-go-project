package dao

import (
	"context"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
)

// DAO is data access object for users table.
type DAO struct {
	conn *pgxpool.Pool
}

func New(conn *pgxpool.Pool) *DAO {
	return &DAO{conn: conn}
}

func (d *DAO) GetByID(ctx context.Context, userID int) (*User, error) {
	stmt := `SELECT id, name FROM users WHERE id = $1`
	rows, err := d.conn.Query(ctx, stmt, userID)
	if err != nil {
		return nil, err
	}

	var user User
	if err := pgxscan.ScanOne(&user, rows); err != nil {
		return nil, err
	}

	return &user, nil
}
