package main

import (
	"log"

	"github.com/gin-gonic/gin"
	//"github.com/AggressiveGas/ChemTrack/backend/routes" // #TODO: change this to right repo
	"github.com/ekjyotshinh/test/backend/routes" // #TODO: change this to right repo

)

func main() {
	// Initialize Firestore
	routes.InitFirestore()

	// Create a Gin router
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router)

	// Start the server on port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
