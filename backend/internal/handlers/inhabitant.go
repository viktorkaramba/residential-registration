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

	var input entity.InputInhabitant
	if err := c.BindJSON(&input); err != nil {
		return

	}

	inhabitant, err := h.Services.Inhabitant.AddInhabitant(buildingID, input)
	if err != nil {
		return
	}

	token, err := h.Services.Token.GenerateToken(inhabitant.ID)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":    inhabitant.ID,
		"token": token,
	})
}
