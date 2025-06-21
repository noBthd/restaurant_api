package models

type Reservation struct {
	ID                int    `json:"id"`
	TableID           int    `json:"table_id"`
	UserID            int    `json:"user_id"`
	StartTime         string `json:"start_time"`
	EndTime           string `json:"end_time"`
	Is_active         bool   `json:"is_active"`
}