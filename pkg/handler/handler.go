package handler

import (
	"github.com/edamiyan/hotel-manager/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
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
