package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/noBthd/restaurant_api.git/internal/models"
	"github.com/noBthd/restaurant_api.git/internal/services"
)

func GetAllReservationsHandler(c *gin.Context) {
	reservations, err := services.GetAllReservations()
	if err != nil {
		log.Printf("Error fetching reservations: %v", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error":   "Failed to fetch reservations",
				"details": err.Error(),
			})
		return
	}

	if len(reservations) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No reservations found"})
		return
	}

	c.JSON(http.StatusOK, reservations)
}

func CreateReservationHandler(c *gin.Context) {
	var reservation models.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	err := services.CreateReservation(reservation)
	if err != nil {
		log.Printf("Error creating reservation: %v", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error":   "Failed to create reservation",
				"details": err.Error(),
			})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Reservation created successfully"})
}

func CancelReservationHandler(c *gin.Context) {
	reservationID := c.Param("id")
	if reservationID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Reservation ID is required"})
		return
	}

	ID, err := strconv.Atoi(reservationID)
	if err != nil {
		log.Printf("Error converting reservation ID to integer: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reservation ID format"})
		return
	}

	err = services.CancelReservation(ID)
	if err != nil {
		log.Printf("Error canceling reservation: %v", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error":   "Failed to cancel reservation",
				"details": err.Error(),
			})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation canceled successfully"})
}

func GetReservedTableByUserIDHandler(c *gin.Context) {
	userIDStr := c.Param("user_id")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		log.Printf("Error converting user ID to integer: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	reservations, err := services.GetReservedTableByUserID(userID)
	if err != nil {
		log.Printf("Error fetching reservations for user ID %d: %v", userID, err)
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error":   "Failed to fetch reservations for user",
				"details": err.Error(),
			})
		return
	}

	if len(reservations) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No reservations found for this user"})
		return
	}

	c.JSON(http.StatusOK, reservations)
}

func GetAllTodayReservationsHandler(c *gin.Context) {
	reservations, err := services.GetAllTodayReservations()
	if err != nil {
		log.Printf("Error fetching today's reservations: %v", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error":   "Failed to fetch today's reservations",
				"details": err.Error(),
			})
		return
	}

	if len(reservations) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No reservations found for today"})
		return
	}

	c.JSON(http.StatusOK, reservations)
}

func GetReservationByDateHandler(c *gin.Context) {
	date := c.Param("date")
	if date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Date is required"})
		return
	}

	reservations, err := services.GetReservationByDate(date)
	if err != nil {
		log.Printf("Error fetching reservations for date %s: %v", date, err)
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error":   "Failed to fetch reservations for the specified date",
				"details": err.Error(),
			})
		return
	}

	if len(reservations) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No reservations found for the specified date"})
		return
	}

	c.JSON(http.StatusOK, reservations)
}

func GetAllFreeSlotToReserveHandler(c *gin.Context) {
	date := c.Param("date")
	if date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Date is required as query param"})
		return
	}

	// Define the time intervals
	intervals := []struct {
		Start string
		End   string
		Label string
	}{
		{"10:00", "11:00", "10:00-11:00"},
		{"12:00", "13:00", "12:00-13:00"},
		{"14:00", "15:00", "14:00-15:00"},
		{"16:00", "17:00", "16:00-17:00"},
		{"18:00", "20:00", "18:00-20:00"},
		{"21:00", "22:00", "21:00-22:00"},
	}

	// Get all tables
	tables, err := services.GetAllTables()
	if err != nil {
		log.Printf("Error fetching tables: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tables"})
		return
	}

	// Get all reservations for the date
	reservations, err := services.GetReservationByDate(date)
	if err != nil {
		log.Printf("Error fetching reservations for date %s: %v", date, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reservations for the specified date"})
		return
	}

	// Build a map: tableID -> interval label -> occupied/free
	result := []gin.H{}
	for _, table := range tables {
		row := gin.H{"TableID": table.ID}
		for _, interval := range intervals {
			occupied := false
			// Parse interval start and end as time.Time
			intervalStart, _ := time.Parse("15:04", interval.Start)
			intervalEnd, _ := time.Parse("15:04", interval.End)
			for _, res := range reservations {
				// Parse reservation start and end as time.Time (extract only time part)
				resStart, err1 := time.Parse(time.RFC3339, res.StartTime)
				resEnd, err2 := time.Parse(time.RFC3339, res.EndTime)
				if err1 != nil || err2 != nil {
					continue
				}
				// Check if reservation is for this table and overlaps with interval
				if res.TableID == table.ID {
					// Get only the time part for comparison
					resStartTime := resStart.Format("15:04")
					resEndTime := resEnd.Format("15:04")
					resStartT, _ := time.Parse("15:04", resStartTime)
					resEndT, _ := time.Parse("15:04", resEndTime)
					// Overlap check
					if resStartT.Before(intervalEnd) && resEndT.After(intervalStart) {
						occupied = true
						break
					}
				}
			}
			if occupied {
				row[interval.Label] = "Occupied"
			} else {
				row[interval.Label] = "Free"
			}
		}
		result = append(result, row)
	}

	c.JSON(http.StatusOK, result)
}