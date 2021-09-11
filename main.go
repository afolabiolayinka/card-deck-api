package main

import "github.com/gin-gonic/gin"

import "api.card-deck/views"


// app entry

func main () {
	// gin router instance

	router := gin.Default()

  router.POST("/create-deck", views.CreateDeck) // `CreateDeck` endpoint
  router.GET("/open-deck/:deck_id", views.OpenDeck) // `OpenDeck` endpoint
  router.GET("/draw-card", views.DrawCard) // `DrawCard` endoint

  router.Run("localhost:8080") // server router on port 8080
}
