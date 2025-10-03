package db

import (
	"context"
	"database/sql"
)

func CreateWineTag(db *sql.DB, ctx context.Context, wineID int64, tagID int64) (*WineTag, error) {
	wineTag := &WineTag{}
	err := db.QueryRowContext(ctx,
		`INSERT INTO wine_tags (wine_id, tag_id) VALUES ($1, $2) RETURNING id, wine_id, tag_id`,
		wineID, tagID,
	).Scan(&wineTag.ID, &wineTag.WineID, &wineTag.TagID)
	if err != nil {
		return nil, err
	}
	return wineTag, nil
}

func DeleteWineTag(db *sql.DB, ctx context.Context, wineID int64, tagID int64) error {
	_, err := db.ExecContext(ctx, `DELETE FROM wine_tags WHERE wine_id = $1 AND tag_id = $2`, wineID, tagID)
	return err
}
