package routes

import (
	"benchmark/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	// USER ROUTES

	router.POST(
		"/users",
		controllers.CreateUser,
	)

	router.GET(
		"/users",
		controllers.GetUsers,
	)

	router.GET(
		"/users/:id",
		controllers.GetUserByID,
	)

	// EVENT ROUTES

	router.POST(
		"/events",
		controllers.CreateEvent,
	)

	router.GET(
		"/events/user/:id",
		controllers.GetEventsByUser,
	)

	router.GET(
		"/events/recent",
		controllers.GetRecentEvents,
	)
}