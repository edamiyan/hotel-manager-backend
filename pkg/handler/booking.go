package handler

import (
	hotelManager "github.com/edamiyan/hotel-manager"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type getAllBookingsResponse struct {
	Data []hotelManager.Booking `json:"data"`
}

type getAllUserBookingsResponse struct {
	Data []hotelManager.BookingTimeline `json:"data"`
}


func (h *Handler) createBooking(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	roomId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid booking id param")
		return
	}

	var input hotelManager.Booking
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Booking.Create(userId, roomId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllBookings(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	roomId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid booking id param")
		return
	}

	bookings, err := h.services.Booking.GetAll(userId, roomId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBookingsResponse{
		Data: bookings,
	})
}

func (h *Handler) getBookingById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	bookingId, err := strconv.Atoi(c.Param("booking_id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "invalid booking id param")
		return
	}

	booking, err := h.services.Booking.GetById(userId, bookingId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, booking)
}

func (h *Handler) updateBooking(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("booking_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	var input hotelManager.UpdateBookingInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Booking.Update(userId, id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) deleteBooking(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("booking_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	err = h.services.Booking.Delete(userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}

func (h *Handler) getRoomIdByBooking(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	bookingId, err := strconv.Atoi(c.Param("booking_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	roomId, err := h.services.Booking.GetRoomIdByBooking(userId, bookingId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"room_id": roomId,
	})
}

func (h *Handler) getAllUserBookings(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}


	bookings, err := h.services.Booking.GetAllUserBookings(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllUserBookingsResponse{
		bookings,
	})
}
