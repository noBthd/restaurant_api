package main

import (
	"github.com/gin-gonic/gin"
	"github.com/noBthd/restaurant_api.git/internal/config"
	"github.com/noBthd/restaurant_api.git/internal/db"
	"github.com/noBthd/restaurant_api.git/internal/handlers"
)

func main() {
	// Load configuration
	db.ConnectDB(config.GetConfig())

	// Initialize the Gin router
	router := gin.Default()

	// Define routes
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

	//=========================================
	// USER ROUTES
	router.GET("/user", handlers.GetUserByEmailHandler)
	// auth/login
	router.POST("/auth/login", handlers.LoginUserHandler)
	// auth/register
	router.POST("/auth/register", handlers.CreateUserHandler)
	
	//=========================================
	// RESERVATION ROUTES
	router.GET("/reservations", handlers.GetAllReservationsHandler)
	// Get reservation by user ID
	router.GET("/reservations/user/:user_id", handlers.GetReservedTableByUserIDHandler)
	// Create a reservation
	router.POST("/reservations/create", handlers.CreateReservationHandler)
	// Cancel a reservation
	router.PATCH("/reservations/cancel/:id", handlers.CancelReservationHandler)
	// Get reservation by date
	router.GET("/reservations/date/:date", handlers.GetReservationByDateHandler)

	//=========================================
	// MENU ROUTES
	router.GET("/menu", handlers.GetAllMenuItemsHandler)

	//=========================================
	// BILL ROUTES
	router.GET("/bills", handlers.GetAllBillsHandler)
	// Get bill by ReservationID
	router.GET("/bills/:reservation_id", handlers.GetBillByReservationIDHandler)
	// Pay a bill
	router.PATCH("/bills/pay/:reservation_id", handlers.PayBillHandler)
	//

	//==========================================
	// ORDER ROUTES
	router.GET("/orders", handlers.GetAllMenuOrdersHandler)
	// Get order by OrderID
	router.GET("/orders/:id", handlers.GetMenuOrderByIDHandler)
	// Make an order
	router.POST("/orders/create", handlers.CreateMenuOrderHandler)

	//==========================================
	// UTILITIES ROUTES
	// Adding served table

	//===========================================
	// ADMIN ROUTES

	// Adding waiter
	router.POST("/admin/waiter/create", handlers.CreateWaiterHandler)
	// Adding shift
	router.POST("/admin/shift/create/table/:id", handlers.CreateShiftHandler)
	// Adding served table
	router.PATCH("/admin/shift/add_served_table/:shift_id", handlers.AddServedTableHandler)
	// Remove user
	router.DELETE("/admin/user/remove/:id", handlers.RemoveUserHandler)
	// Make user admin
	router.PATCH("/admin/user/make_admin/:id", handlers.MakeUserAdminHandler)
	// Get all orders by reservationID
	router.GET("/admin/menu_orders/reservation/:id", handlers.GetAllMenuOrdersByReservationIDHandler)
	// Adding served table by reservation ID
	router.PATCH("/admin/shift/add_served_table_by_reservation/:reservation_id", handlers.AddServedTableByReservationHandler)
	// Getting all today reservations
	router.GET("/admin/reservations/today", handlers.GetAllTodayReservationsHandler)


	router.Run(":8080")	
}
