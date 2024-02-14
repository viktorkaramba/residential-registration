package handlers

import (
	"net/http"
	"residential-registration/backend/internal/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) registerInhabitant(c *gin.Context) {

	buildingID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return
	}

	var input entity.InputUser
	if err := c.BindJSON(&input); err != nil {
		return

	}

	User, err := h.Services.User.AddUser(buildingID, input)
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
