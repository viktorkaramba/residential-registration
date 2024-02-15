package handlers

import (
	"bytes"
	"io"
	"net/http"
	"residential-registration/backend/internal/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) registerInhabitant(c *gin.Context) {

	osbbID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return
	}

	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	// Check if there are any additional fields in the JSON body
	if err := h.validateJSONTags(body, entity.InputUser{}); err != nil {
		return
	}

	var input entity.InputUser
	if err := c.BindJSON(&input); err != nil {
		return

	}

	User, err := h.Services.User.AddUser(osbbID, input)
	if err != nil {
		return
	}

	token, err := h.Services.Token.GenerateToken(User.ID)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":    User.ID,
		"token": token,
	})
}
