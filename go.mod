module api/card-deck-api

go 1.16

require (
	api.card-deck/models v1.0.0
	api.card-deck/views v1.0.0
	api.card-deck/router v1.0.0
	api.card-deck/utils v1.0.0
)

replace (
	api.card-deck/models v1.0.0 => "./models"
	api.card-deck/views v1.0.0 => "./views"
	api.card-deck/router v1.0.0 => "./router"
	api.card-deck/utils v1.0.0 => "./utils"
)
