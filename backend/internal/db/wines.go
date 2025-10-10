package db

import (
	"context"
	"database/sql"
)

func CreateWine(db *sql.DB, ctx context.Context, wine *Wine) (*Wine, error) {
	row := db.QueryRowContext(ctx, `INSERT INTO wines (user_id, name, wine_producer, region, stock_quantity, storage_location_id, photo_url) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		wine.UserID, wine.Name, wine.WineProducer, wine.Region, wine.StockQuantity, wine.StorageLocationID, wine.PhotoURL)
	if err := row.Scan(&wine.ID); err != nil {
		return nil, err
	}
	return wine, nil
}

func GetWine(db *sql.DB, ctx context.Context, id int64) (*Wine, error) {
	wine := &Wine{}
	row := db.QueryRowContext(ctx, `SELECT id, user_id, name, wine_producer, region, stock_quantity, storage_location_id, photo_url FROM wines WHERE id = $1`, id)
	if err := row.Scan(&wine.ID, &wine.UserID, &wine.Name, &wine.WineProducer, &wine.Region, &wine.StockQuantity, &wine.StorageLocationID, &wine.PhotoURL); err != nil {
		return nil, err
	}
	return wine, nil
}

func DeleteWine(db *sql.DB, ctx context.Context, id int64) error {
	_, err := db.ExecContext(ctx, `DELETE FROM wines WHERE id = $1`, id)
	return err
}

func UpdateWine(db *sql.DB, ctx context.Context, wine *Wine) error {
	_, err := db.ExecContext(ctx, `UPDATE wines SET name = $1, wine_producer = $2, region = $3, stock_quantity = $4, storage_location_id = $5, photo_url = $6 WHERE id = $7`,
		wine.Name, wine.WineProducer, wine.Region, wine.StockQuantity, wine.StorageLocationID, wine.PhotoURL, wine.ID)
	return err
}

func ListWinesByUser(db *sql.DB, ctx context.Context, userID int64) ([]*Wine, error) {
	rows, err := db.QueryContext(ctx, `SELECT id, user_id, name, wine_producer, region, stock_quantity, storage_location_id, photo_url FROM wines WHERE user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var wines []*Wine
	for rows.Next() {
		wine := &Wine{}
		if err := rows.Scan(&wine.ID, &wine.UserID, &wine.Name, &wine.WineProducer, &wine.Region, &wine.StockQuantity, &wine.StorageLocationID, &wine.PhotoURL); err != nil {
			return nil, err
		}
		wines = append(wines, wine)
	}
	return wines, nil
}
