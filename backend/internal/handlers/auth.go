package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"residential-registration/backend/internal/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) registerInhabitant(c *gin.Context) {
	logger := h.Logger.Named("registerInhabitant").WithContext(c)

	osbbID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		logger.Error("failed to parse osbb id", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to parse osbb id: %w", err))
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error("failed to read body request", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to read body request: %w", err))
		return
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	// Check if there are any additional fields in the JSON body
	if err := h.validateJSONTags(body, entity.EventUserPayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventUserPayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to bind JSON: %w", err))
		return
	}

	inhabitant, err := h.Services.Auth.AddUser(osbbID, input)
	if err != nil {
		logger.Error("failed to add inhabitant", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to add inhabitant: %w", err))
		return
	}

	token, err := h.Services.Token.GenerateToken(inhabitant.ID)
	if err != nil {
		logger.Error("failed to generate token", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to generate token: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id":    inhabitant.ID,
		"token": token,
	})
}

func (h *Handler) login(c *gin.Context) {
	logger := h.Logger.Named("login").WithContext(c)

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error("failed to read body request", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to read body request: %w", err))
		return
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	// Check if there are any additional fields in the JSON body
	if err := h.validateJSONTags(body, entity.EventLoginPayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventLoginPayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to bind JSON: %w", err))
		return
	}

	inhabitant, err := h.Services.Auth.Login(input)
	if err != nil {
		logger.Error("failed to login", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to login: %w", err))
		return
	}

	token, err := h.Services.Token.GenerateToken(inhabitant.ID)
	if err != nil {
		logger.Error("failed to generate token", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to generate token: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) logout(c *gin.Context) {
	logger := h.Logger.Named("logout").WithContext(c)

	token, err := h.checkHeaderToken(c)
	if err != nil {
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed check header token: %w", err))
		return
	}

	err = h.Services.Auth.Logout(entity.TokenValue(token))
	if err != nil {
		logger.Error("failed to logout", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to logout: %w", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Logout successfully",
	})
}
