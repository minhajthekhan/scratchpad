package legos

import (
	"context"
	"database/sql"
)

type repository struct {
	db sql.DB
}

func (r *repository) GetLogoBySpecification(ctx context.Context, spec Specification[Lego]) ([]string, error) {
	query, args := spec.AsSQL()
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	positions := make([]string, 0)
	for rows.Next() {
		var shelfPosition string
		if err := rows.Scan(&shelfPosition); err != nil {
			return nil, err
		}
		positions = append(positions, shelfPosition)
	}
	return positions, nil
}

func (r *repository) GetLegoShelfPositions(ctx context.Context, d1, d2 LegoDimension, colors []string, limit int) ([]string, error) {
	args := []any{colors, d1.Size, d1.Height, d2.Height, d2.Size, limit}
	rows, err := r.db.Query(`
		SELECT shelf_position
		FROM legos
		WHERE color = ANY($1) 
		AND ((size = $2 AND height < $3) OR (size = $4 AND height > $5)) LIMIT $6
	`, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	return r.scanRowsAndGetPositions(rows)
}

func (r *repository) scanRowsAndGetPositions(rows *sql.Rows) ([]string, error) {
	positions := make([]string, 0)
	for rows.Next() {
		var shelfPosition string
		if err := rows.Scan(&shelfPosition); err != nil {
			return nil, err
		}
		positions = append(positions, shelfPosition)
	}
	return positions, nil
}
