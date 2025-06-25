package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		c.Abort()
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header format")
		c.Abort()
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error()) // можно отдать текст ошибки
		c.Abort()
		return
	}

	c.Set("userCtx", userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get("userCtx")
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "No user context")
		return 0, errors.New("No user context")
	}
	idInt, ok := id.(int)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "No user id in context")
		return 0, errors.New("No user id in context")
	}
	return idInt, nil
}
