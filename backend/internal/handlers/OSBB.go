package handlers

import (
	"fmt"
	"net/http"
	"residential-registration/backend/internal/entity"

	"github.com/gin-gonic/gin"
)

func (h *Handler) registerOSBB(c *gin.Context) {

	var input entity.InputOSBB
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(input)
	osbb, err := h.Services.OSBB.AddOSBB(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	token, err := h.Services.Token.GenerateToken(osbb.OSBBHead.ID)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": osbb.OSBBHead.ID,
		"token":   token,
	})
}
