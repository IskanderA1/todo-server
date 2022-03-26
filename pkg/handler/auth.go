package handler

import (
	"net/http"

	"github.com/IskanderA1/todo-server"
	"github.com/gin-gonic/gin"
)

func (h *Handler) singUp(c *gin.Context) {
	var input todo.User
	println("singUp")
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.service.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) singIn(c *gin.Context) {

}
