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

func (h *Handler) registerOSBB(c *gin.Context) {
	logger := h.Logger.Named("registerOSBB").WithContext(c)

	var input entity.EventOSBBPayload

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error("failed to rea-d body request", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to read body request: %w", err))
		return
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	// Check if there are any additional fields in the JSON body
	if err := h.validateJSONTags(body, entity.EventOSBBPayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bing JSON", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to bing JSON: %w", err))
		return
	}

	osbb, err := h.Services.OSBB.AddOSBB(input)
	if err != nil {
		logger.Error("failed to add osbb", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to add osbb: %w", err))
		return
	}

	token, err := h.Services.Token.GenerateToken(osbb.OSBBHead.ID)
	if err != nil {
		logger.Error("failed to generate token", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to generate token: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user_id": osbb.OSBBHead.ID,
		"token":   token,
	})
}

func (h *Handler) addAnnouncement(c *gin.Context) {
	logger := h.Logger.Named("addAnnouncement").WithContext(c)

	userID, err := h.getUserId(c)
	if err != nil {
		logger.Error("failed to get user id", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get user id: %w", err))
		return
	}

	osbbID, err := strconv.ParseUint(c.Param("osbbID"), 10, 64)
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
	if err := h.validateJSONTags(body, entity.EventAnnouncementPayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventAnnouncementPayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to bind JSON: %w", err))
		return
	}

	announcement, err := h.Services.OSBB.AddAnnouncement(userID, osbbID, input)
	if err != nil {
		logger.Error("failed to add announcement", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to add announcement: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": announcement.ID,
	})
}

func (h *Handler) addPoll(c *gin.Context) {
	logger := h.Logger.Named("addPoll").WithContext(c)

	userID, err := h.getUserId(c)
	if err != nil {
		logger.Error("failed to get user id", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get user id: %w", err))
		return
	}

	osbbID, err := strconv.ParseUint(c.Param("osbbID"), 10, 64)
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
	if err := h.validateJSONTags(body, entity.EventPollPayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventPollPayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to bind JSON: %w", err))
		return
	}

	poll, err := h.Services.OSBB.AddPoll(userID, osbbID, input)
	if err != nil {
		logger.Error("failed to add poll", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to add poll: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": poll.ID,
	})
}

func (h *Handler) addPollTest(c *gin.Context) {
	logger := h.Logger.Named("addPollTest").WithContext(c)

	userID, err := h.getUserId(c)
	if err != nil {
		logger.Error("failed to get user id", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get user id: %w", err))
		return
	}

	osbbID, err := strconv.ParseUint(c.Param("osbbID"), 10, 64)
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
	if err := h.validateJSONTags(body, entity.EventPollTestPayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventPollTestPayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to bind JSON: %w", err))
		return
	}

	poll, err := h.Services.OSBB.AddPollTest(userID, osbbID, input)
	if err != nil {
		logger.Error("failed to add poll test", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to add poll test: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": poll.ID,
	})
}

func (h *Handler) addPollAnswer(c *gin.Context) {
	logger := h.Logger.Named("addPollAnswer").WithContext(c)

	userID, err := h.getUserId(c)
	if err != nil {
		logger.Error("failed to get user id", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get user id: %w", err))
		return
	}

	pollID, err := strconv.ParseUint(c.Param("pollID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse poll id", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to parse poll id: %w", err))
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
	if err := h.validateJSONTags(body, entity.EventPollAnswerPayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventPollAnswerPayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to bind JSON: %w", err))
		return
	}

	answer, err := h.Services.OSBB.AddPollAnswer(userID, pollID, input)
	if err != nil {
		logger.Error("failed to add poll answer", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to add poll answer: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": answer.ID,
	})
}

func (h *Handler) addPollAnswerTest(c *gin.Context) {
	logger := h.Logger.Named("addPollAnswerTest").WithContext(c)

	userID, err := h.getUserId(c)
	if err != nil {
		logger.Error("failed to get user id", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get user id: %w", err))
		return
	}

	pollID, err := strconv.ParseUint(c.Param("pollID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse poll id", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to parse poll id: %w", err))
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
	if err := h.validateJSONTags(body, entity.EventPollAnswerTestPayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventPollAnswerTestPayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to bind JSON: %w", err))
		return
	}

	poll, err := h.Services.OSBB.AddPollAnswerTest(userID, pollID, input)
	if err != nil {
		logger.Error("failed to add poll answer test", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to add poll answer test: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": poll.ID,
	})
}

func (h *Handler) addPayment(c *gin.Context) {
	logger := h.Logger.Named("addPayment").WithContext(c)

	userID, err := h.getUserId(c)
	if err != nil {
		logger.Error("failed to get user id", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get user id: %w", err))
		return
	}

	osbbID, err := strconv.ParseUint(c.Param("osbbID"), 10, 64)
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
	if err := h.validateJSONTags(body, entity.EventPaymentPayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventPaymentPayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to bind JSON: %w", err))
		return
	}

	payment, err := h.Services.OSBB.AddPayment(userID, osbbID, input)
	if err != nil {
		logger.Error("failed to add payment", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to add payment: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": payment.ID,
	})
}

func (h *Handler) makePurchase(c *gin.Context) {
	logger := h.Logger.Named("makePayment").WithContext(c)

	userID, err := h.getUserId(c)
	if err != nil {
		logger.Error("failed to get user id", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get user id: %w", err))
		return
	}

	paymentID, err := strconv.ParseUint(c.Param("paymentID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse osbb id", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to parse osbb id: %w", err))
		return
	}

	userPayment, err := h.Services.OSBB.AddPurchase(userID, paymentID)
	if err != nil {
		logger.Error("failed to add user payment", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to add user payment: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": userPayment.ID,
	})
}
