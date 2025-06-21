package handler

import (
	golang_todo_app "github.com/Numbone/golang-todo-app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input golang_todo_app.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	h.services
}

func (h *Handler) signIn(c *gin.Context) {

}
