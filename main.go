package main

import "api.card-deck/controllers"


// app entry

func main () {
	// run server

  controllers.SetupRouter().Run("localhost:8080")
}
