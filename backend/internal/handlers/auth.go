package handlers

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"residential-registration/backend/internal/entity"
	"residential-registration/backend/pkg/errs"

	"github.com/gin-gonic/gin"
)

func (h *Handler) login(c *gin.Context) {
	logger := h.Logger.Named("login").WithContext(c)

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error("failed to read body request", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(errors.New("failed to read body request")).Code("Failed body validation").Kind(errs.Validation))
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
		h.sendErrResponse(c, h.Logger,
			errs.Err(errors.New("failed to bind JSON")).Code("Failed bind JSON").Kind(errs.Validation))
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

func (h *Handler) refreshToken(c *gin.Context) {
	logger := h.Logger.Named("refreshToken").WithContext(c)

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error("failed to read body request", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(errors.New("failed to read body request")).Code("Failed body validation").Kind(errs.Validation))

		return
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	// Check if there are any additional fields in the JSON body
	if err := h.validateJSONTags(body, entity.EventTokenPayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventTokenPayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(errors.New("failed to bind JSON")).Code("Failed bind JSON").Kind(errs.Validation))
		return
	}

	oldToken, err := h.Services.Token.GetByToken(string(input.TokenValue))
	if err != nil {
		logger.Error("failed to get old token", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get old token: %w", err))
		return
	}
	if oldToken.Revoked {
		logger.Error("failed to refresh token", "error", errors.New("token is revoked"))
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to refresh token: %w", errors.New("token is revoked")))
		return
	}
	newToken, err := h.Services.Token.RefreshToken(oldToken.UserID)
	if err != nil {
		logger.Error("failed to refresh token", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to refresh token: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": newToken,
	})
}
