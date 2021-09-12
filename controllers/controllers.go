package controllers

import "github.com/gin-gonic/gin"

import "api.card-deck/views"


// mode -> developement (true) / production (false)

const DEBUG bool = false


func SetupRouter () *gin.Engine {
	// set DEBUG mode to production / developement

	if !DEBUG {
		gin.SetMode(gin.ReleaseMode)
	}

	// gin router instance

	routes := gin.Default()

	routes.POST("/create-deck", views.CreateDeck) // `CreateDeck` endpoint
	routes.GET("/open-deck/:deck_id", views.OpenDeck) // `OpenDeck` endpoint
	routes.GET("/draw-card", views.DrawCard) // `DrawCard` endoint

	return routes
}
