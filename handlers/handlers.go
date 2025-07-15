package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Flights []Flight

func AddFlight(c *gin.Context) {
	var input struct {
		Src      string `json:"src"`
		Dst      string `json:"dst"`
		MaxLimit int    `json:"max_limit"`
	}

	if err := c.ShouldBindJSON(&input); err != nil || input.Src == "" || input.Dst == "" || input.MaxLimit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input, 3no field chahiye limit badi zero se",
		})
		return
	}

	newFlight := NewFlight(input.Src, input.Dst, input.MaxLimit)
	Flights = append(Flights, *newFlight)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Flight added successfully",
		"flight":  newFlight,
	})
}

func GetAllFlights(c *gin.Context) {
	if Flights == nil {
		Flights = []Flight{}
	}
	c.JSON(http.StatusOK, gin.H{
		"flights": Flights,
	})
}

func GetFlightByID(c *gin.Context) {
	id := c.Param("id")

	for _, flight := range Flights {
		if flight.Id == id {
			c.JSON(http.StatusOK, gin.H{
				"flight": flight,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Flight not found",
	})
}

func UpdateFlight(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Src string `json:"src"`
		Dst string `json:"dst"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	for i := range Flights {
		if Flights[i].Id == id {
			if input.Src != "" {
				Flights[i].Src = input.Src
			}
			if input.Dst != "" {
				Flights[i].Dst = input.Dst
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "Flight updated successfully",
				"flight":  Flights[i],
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{
		"error": "Flight not found"})
}

func DeleteFlight(c *gin.Context) {
	id := c.Param("id")

	for i, flight := range Flights {
		if flight.Id == id {
			Flights = append(Flights[:i], Flights[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Flight deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
}

// func addpassengers(){
//  input me user se num of passengers lega
//  passengers += input
//  par only if add krne ke baad max_limit us flight ki reach na ho tabhi add krna warna bad request usme likhna ki limit reached try lower number of passengers
// }
// POST /:id/passengers
// {
// 	"passengers": 5
// }

func AddPassengers(c *gin.Context) {
	id := c.Param("id")

	var input struct {
		Passengers int `json:"passengers"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	for i := range Flights {
		if Flights[i].Id == id {
			if Flights[i].Passengers+input.Passengers > Flights[i].MaxLimit {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Aircraft Full",
				})
				return
			}

			Flights[i].Passengers += input.Passengers

			c.JSON(http.StatusOK, gin.H{
				"message":      "Passengers added successfully",
				"total_Passengers":    Flights[i].Passengers,
				"max_capacity": Flights[i].MaxLimit,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Flight not found"})
}
