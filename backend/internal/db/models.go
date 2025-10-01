package db

type User struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type StorageLocation struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	Name   string `json:"location_name"`
}

type Tag struct {
	ID     int64  `json:"id"`
	UserID int64  `json:"user_id"`
	WineID int64  `json:"wine_id"`
	Name   string `json:"tag_name"`
	Color  string `json:"color"`
}

type Wine struct {
	ID                int64   `json:"id"`
	UserID            int64   `json:"user_id"`
	Name              string  `json:"name"`
	WineProducer      string  `json:"wine_producer"`
	Region            string  `json:"region"`
	StockQuantity     int     `json:"stock_quantity"`
	StorageLocationID *int64  `json:"storage_location_id"`
	PhotoURL          *string `json:"photo_url"`
}
