package main

import "github.com/gin-gonic/gin"

import . "api.card-deck/views"


// app entry

func main () {
	// gin router instance

	router := gin.Default()

  router.POST("/create-deck", CreateDeck) // `CreateDeck` endpoint
  router.GET("/open-deck/:deck_id", OpenDeck) // `OpenDeck` endpoint
  router.GET("/draw-card", DrawCard) // `DrawCard` endoint

  router.Run("localhost:8080") // server router on port 8080
}
