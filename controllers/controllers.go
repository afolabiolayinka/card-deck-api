package controllers

import "github.com/gin-gonic/gin"

import "api.card-deck/views"


func InitRouter () {
	// gin router instance

	routes := gin.Default()

	routes.POST("/create-deck", views.CreateDeck) // `CreateDeck` endpoint
	routes.GET("/open-deck/:deck_id", views.OpenDeck) // `OpenDeck` endpoint
	routes.GET("/draw-card", views.DrawCard) // `DrawCard` endoint

	routes.Run("localhost:8080") // server on port 8080
}
