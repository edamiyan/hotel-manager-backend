package handler

import (
	hotelManager "github.com/edamiyan/hotel-manager"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createRoom(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input hotelManager.Room
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Room.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllRoomsResponse struct {
	Data []hotelManager.Room `json:"data"`
}

func (h *Handler) getAllRooms(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	rooms, err := h.services.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllRoomsResponse{
		Data: rooms,
	})
}

func (h *Handler) getRoomById(c *gin.Context) {

}

func (h *Handler) updateRoom(c *gin.Context) {

}

func (h *Handler) deleteRoom(c *gin.Context) {

}
