package models

type MenuOrder struct {
	ID 			int    `json:"id"`
	OrderID 	int    `json:"order_id"`
	MenuItemID 	int    `json:"menu_item_id"`
	Quantity 	int    `json:"quantity"`
	Price   	int    `json:"price"`
}