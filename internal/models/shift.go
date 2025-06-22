package models

type Shift struct {
	ID        	 int    `json:"id"`
	WaiterID  	 int    `json:"waiter_id"`
	Date 	  	 string `json:"date"`
	TablesServed int    `json:"tables_served"`
}