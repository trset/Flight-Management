package routes

import (
	"flight_management/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/", handlers.AddFlight)
	r.GET("/", handlers.GetAllFlights)
	r.GET("/:id", handlers.GetFlightByID)
	r.PUT("/:id", handlers.UpdateFlight)
	r.PUT("/:id/passengers", handlers.AddPassengers)
	r.DELETE("/:id", handlers.DeleteFlight)
	

}
