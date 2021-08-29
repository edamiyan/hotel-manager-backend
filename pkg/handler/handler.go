package handler

import (
	"github.com/edamiyan/hotel-manager/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, GET, DELETE, PUT, OPTIONS")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		rooms := api.Group("/rooms")
		{
			rooms.POST("/", h.createRoom)
			rooms.GET("/", h.getAllRooms)
			rooms.GET("/:id", h.getRoomById)
			rooms.PUT("/:id", h.updateRoom)
			rooms.DELETE("/:id", h.deleteRoom)

			bookings := rooms.Group(":id/bookings")
			{
				bookings.POST("/", h.createBooking)
				bookings.GET("/", h.getAllBookings)
				bookings.GET("/:booking_id", h.getBookingById)
				bookings.PUT("/:booking_id", h.updateBooking)
				bookings.DELETE("/:booking_id", h.deleteBooking)
			}
		}
	}

	return router
}
