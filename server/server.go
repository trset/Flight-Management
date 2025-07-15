package server

import (
	"flight_management/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	// API routes
	routes.RegisterRoutes(r)

	// ✅ Serve frontend at a different route to avoid conflict
	r.StaticFile("/home", "./index.html")
	r.StaticFile("/manifest.json", "./manifest.json")
	r.StaticFile("/service-worker.js", "./service-worker.js")
	r.StaticFile("/icon-192.png", "./icon-192.png")
	r.StaticFile("/icon-512.png", "./icon-512.png")

	// Start server
	if err := r.Run(":8080"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
