package handlers

import (
	"fmt"
	"net/http"
	"residential-registration/backend/internal/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) registerOSBB(c *gin.Context) {
	var input entity.InputOSBB
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		return
	}

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

func (h *Handler) addAnnouncement(c *gin.Context) {

	userID, err := getUserId(c)
	if err != nil {
		fmt.Println(err)
		return
	}

	osbbID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return
	}

	var input entity.InputAnnouncement
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		return
	}

	announcement, err := h.Services.OSBB.AddAnnouncement(userID, osbbID, input)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": announcement.ID,
	})
}
