package models

type Bill struct {
	ID          	int     `json:"id"`
	ReservationID 	int 	`json:"reservation_id"`
	TotalPrice 		int     `json:"total_price"`
	IsPaid     		bool    `json:"is_paid"`
}