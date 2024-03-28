package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"residential-registration/backend/internal/entity"
	"residential-registration/backend/pkg/errs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) registerInhabitant(c *gin.Context) {
	logger := h.Logger.Named("registerInhabitant").WithContext(c)

	osbbID, err := strconv.ParseUint(c.Param("osbbID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse osbb id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to parse osbb id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
		return
	}

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error("failed to read body request", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to read body request: %w", err)).Code("Failed body validation").Kind(errs.Validation))
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
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to bind JSON: %w", err)).Code("Failed bind JSON").Kind(errs.Validation))
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

func (h *Handler) getAllInhabitants(c *gin.Context) {
	logger := h.Logger.Named("getAllInhabitans").WithContext(c)

	userID, err := h.getUserId(c)
	if err != nil {
		logger.Error("failed to get user id", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get user id: %w", err))
		return
	}

	osbbID, err := strconv.ParseUint(c.Param("osbbID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse osbb id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to parse osbb id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
		return
	}

	inhabitants, err := h.Services.OSBB.ListInhabitants(userID, osbbID, services.UserFilter{
		OSBBID:         &osbbID,
		IsApproved:     typecast.ToPtr(true),
		WithIsApproved: typecast.ToPtr(true),
		WithApartment:  typecast.ToPtr(true),
	})
	if err != nil {
		logger.Error("failed to get all inhabitants", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get all inhabitants: %w", err))
		return
	}

	c.JSON(http.StatusOK, inhabitants)
}

func (h *Handler) getWaitApproveInhabitants(c *gin.Context) {
	logger := h.Logger.Named("getWaitApproveInhabitants").WithContext(c)

	userID, err := h.getUserId(c)
	if err != nil {
		logger.Error("failed to get user id", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get user id: %w", err))
		return
	}

	osbbID, err := strconv.ParseUint(c.Param("osbbID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse osbb id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to parse osbb id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
		return
	}

	inhabitants, err := h.Services.OSBB.ListInhabitants(userID, osbbID, services.UserFilter{
		OSBBID:         &osbbID,
		IsApproved:     nil,
		WithIsApproved: typecast.ToPtr(true),
		WithApartment:  typecast.ToPtr(true),
	})
  
	if err != nil {
		logger.Error("failed to get all inhabitants", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get all inhabitants: %w", err))
		return
	}

	c.JSON(http.StatusOK, inhabitants)
}

func (h *Handler) getInhabitantsProfile(c *gin.Context) {
	logger := h.Logger.Named("getInhabitantsProfile").WithContext(c)

	userID, err := h.getUserId(c)
	if err != nil {
		logger.Error("failed to get user id", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get user id: %w", err))
		return
	}

	inhabitant, err := h.Services.OSBB.GetInhabitant(userID)
	if err != nil {
		logger.Error("failed to get all inhabitants", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get all inhabitants: %w", err))
		return
	}

	c.JSON(http.StatusOK, inhabitant)
}

func (h *Handler) updateInhabitant(c *gin.Context) {
	logger := h.Logger.Named("updateInhabitant").WithContext(c)

	userID, err := h.getUserId(c)
	if err != nil {
		logger.Error("failed to get user id", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get user id: %w", err))
		return
	}
	osbbID, err := strconv.ParseUint(c.Param("osbbID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse osbb id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to parse osbb id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
		return
	}
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error("failed to read body request", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to read body request: %w", err)).Code("Failed body validation").Kind(errs.Validation))
		return
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	// Check if there are any additional fields in the JSON body
	if err := h.validateJSONTags(body, entity.EventUserUpdatePayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventUserUpdatePayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to bind JSON: %w", err)).Code("Failed bind JSON").Kind(errs.Validation))
		return
	}

	err = h.Services.OSBB.UpdateInhabitant(userID, osbbID, input)
	if err != nil {
		logger.Error("failed to update inhabitant", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to update inhabitant: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
