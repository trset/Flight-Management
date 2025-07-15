package handlers

import "github.com/google/uuid"

type Flight struct {
	Id         string `json:"id"`
	Src        string `json:"src"`
	Dst        string `json:"dst"`
	Passengers int    `json:"passengers"`
	MaxLimit   int    `json:"max_limit"`
	Status     string `json:"status"`
}

func NewFlight(src string, dst string, limit int) *Flight {
	return &Flight{
		Id:         uuid.NewString(),
		Src:        src,
		Dst:        dst,
		Passengers: 0,
		MaxLimit:   limit,
		Status:     FLIGHT_READY,
	}
}

var flights = []Flight{}

