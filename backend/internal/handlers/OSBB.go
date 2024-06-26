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

func (h *Handler) registerOSBB(c *gin.Context) {
	logger := h.Logger.Named("registerOSBB").WithContext(c)

	var input entity.EventOSBBPayload

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error("failed to read body request", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to read body request: %w", err)).Code("Failed body validation").Kind(errs.Validation))
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
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to bind JSON: %w", err)).Code("Failed bind JSON").Kind(errs.Validation))
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

func (h *Handler) getAllOSBB(c *gin.Context) {
	logger := h.Logger.Named("getAllOSBB").WithContext(c)

	osbbs, err := h.Services.OSBB.ListOSBBS()
	if err != nil {
		logger.Error("failed to get list osbb", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get list osbb: %w", err))
		return
	}

	c.JSON(http.StatusOK, osbbs)
}

func (h *Handler) getOSBBProfile(c *gin.Context) {
	logger := h.Logger.Named("getOSBBProfile").WithContext(c)

	userID, err := h.getUserId(c)
	if err != nil {
		logger.Error("failed to get user id", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get user id: %w", err))
		return
	}

	osbb, err := h.Services.OSBB.GetOSBB(userID)
	if err != nil {
		logger.Error("failed to get osbb profile", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get osbb profile: %w", err))
		return
	}

	c.JSON(http.StatusOK, osbb)
}

func (h *Handler) updateOSBB(c *gin.Context) {
	logger := h.Logger.Named("updateOSBB").WithContext(c)

	userID, err := h.getUserId(c)
	if err != nil {
		logger.Error("failed to get user id", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get user id: %w", err))
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
	if err := h.validateJSONTags(body, entity.EventOSBBUpdatePayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventOSBBUpdatePayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to bind JSON: %w", err)).Code("Failed bind JSON").Kind(errs.Validation))
		return
	}

	err = h.Services.OSBB.UpdateOSBB(userID, input)
	if err != nil {
		logger.Error("failed to update osbb profile", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to update osbb profile: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) addApartment(c *gin.Context) {
	logger := h.Logger.Named("addApartment").WithContext(c)

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
	if err := h.validateJSONTags(body, entity.EventApartmentPayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventApartmentPayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to bind JSON: %w", err)).Code("Failed bind JSON").Kind(errs.Validation))
		return
	}

	announcement, err := h.Services.OSBB.AddApartment(userID, osbbID, input)
	if err != nil {
		logger.Error("failed to add announcement", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to add announcement: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": announcement.ID,
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
	if err := h.validateJSONTags(body, entity.EventAnnouncementPayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventAnnouncementPayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to bind JSON: %w", err)).Code("Failed bind JSON").Kind(errs.Validation))
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

func (h *Handler) getAllAnnouncement(c *gin.Context) {
	logger := h.Logger.Named("getAllAnnouncement").WithContext(c)

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

	announcements, err := h.Services.OSBB.ListAnnouncements(userID, osbbID)
	if err != nil {
		logger.Error("failed to get list announcement", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get list announcement: %w", err))
		return
	}

	c.JSON(http.StatusOK, announcements)
}

func (h *Handler) updateAnnouncement(c *gin.Context) {
	logger := h.Logger.Named("updateAnnouncement").WithContext(c)

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

	announcementID, err := strconv.ParseUint(c.Param("announcementID"), 10, 64)
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
	if err := h.validateJSONTags(body, entity.EventAnnouncementUpdatePayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventAnnouncementUpdatePayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to bind JSON: %w", err)).Code("Failed bind JSON").Kind(errs.Validation))
		return
	}

	err = h.Services.OSBB.UpdateAnnouncement(userID, osbbID, announcementID, input)
	if err != nil {
		logger.Error("failed to update announcement", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to update announcement: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) deleteAnnouncement(c *gin.Context) {
	logger := h.Logger.Named("deleteAnnouncement").WithContext(c)

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

	announcementID, err := strconv.ParseUint(c.Param("announcementID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse osbb id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to parse osbb id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
		return
	}

	err = h.Services.OSBB.DeleteAnnouncement(userID, osbbID, announcementID)
	if err != nil {
		logger.Error("failed to delete announcement", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to delete announcement: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
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
	if err := h.validateJSONTags(body, entity.EventPollPayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventPollPayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to bind JSON: %w", err)).Code("Failed bind JSON").Kind(errs.Validation))
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
	if err := h.validateJSONTags(body, entity.EventPollTestPayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventPollTestPayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to bind JSON: %w", err)).Code("Failed bind JSON").Kind(errs.Validation))
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

func (h *Handler) getAllPolls(c *gin.Context) {
	logger := h.Logger.Named("getAllPolls").WithContext(c)

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

	polls, err := h.Services.OSBB.ListPolls(userID, osbbID)
	if err != nil {
		logger.Error("failed to get list polls", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get list polls: %w", err))
		return
	}

	c.JSON(http.StatusOK, polls)
}

func (h *Handler) updatePoll(c *gin.Context) {
	logger := h.Logger.Named("updatePoll").WithContext(c)

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

	pollID, err := strconv.ParseUint(c.Param("pollID"), 10, 64)
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
	if err := h.validateJSONTags(body, entity.EventPollUpdatePayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventPollUpdatePayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to bind JSON: %w", err)).Code("Failed bind JSON").Kind(errs.Validation))
		return
	}

	err = h.Services.OSBB.UpdatePoll(userID, osbbID, pollID, input)
	if err != nil {
		logger.Error("failed to update poll", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to update poll: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) deletePoll(c *gin.Context) {
	logger := h.Logger.Named("deletePoll").WithContext(c)

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

	pollID, err := strconv.ParseUint(c.Param("pollID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse osbb id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to parse osbb id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
		return
	}

	err = h.Services.OSBB.DeletePoll(userID, osbbID, pollID)
	if err != nil {
		logger.Error("failed to delete poll", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to delete poll: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
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

	osbbID, err := strconv.ParseUint(c.Param("osbbID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse osbb id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to parse osbb id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
		return
	}

	pollID, err := strconv.ParseUint(c.Param("pollID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse poll id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to parse poll id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
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
	if err := h.validateJSONTags(body, entity.EventPollAnswerPayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventPollAnswerPayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to bind JSON: %w", err)).Code("Failed bind JSON").Kind(errs.Validation))
		return
	}

	answer, err := h.Services.OSBB.AddPollAnswer(userID, pollID, osbbID, input)
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

	osbbID, err := strconv.ParseUint(c.Param("osbbID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse osbb id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to parse osbb id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
		return
	}

	pollID, err := strconv.ParseUint(c.Param("pollID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse poll id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to parse poll id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
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
	if err := h.validateJSONTags(body, entity.EventPollAnswerTestPayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventPollAnswerTestPayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to bind JSON: %w", err)).Code("Failed bind JSON").Kind(errs.Validation))
		return
	}

	poll, err := h.Services.OSBB.AddPollAnswerTest(userID, pollID, osbbID, input)
	if err != nil {
		logger.Error("failed to add poll answer test", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to add poll answer test: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": poll.ID,
	})
}

func (h *Handler) getUserAnswers(c *gin.Context) {
	logger := h.Logger.Named("getUserAnswers").WithContext(c)

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

	pollID, err := strconv.ParseUint(c.Param("pollID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse osbb id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to parse poll id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
		return
	}

	userAnswer, err := h.Services.OSBB.GetUserAnswer(userID, osbbID, pollID)
	if err != nil {
		logger.Error("failed to get user answers", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get user answers: %w", err))
		return
	}
	c.JSON(http.StatusOK, userAnswer)
}

func (h *Handler) updateUserAnswer(c *gin.Context) {
	logger := h.Logger.Named("updateUserAnswer").WithContext(c)

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

	pollID, err := strconv.ParseUint(c.Param("pollID"), 10, 64)
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
	if err := h.validateJSONTags(body, entity.EventUserAnswerUpdatePayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventUserAnswerUpdatePayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to bind JSON: %w", err)).Code("Failed bind JSON").Kind(errs.Validation))
		return
	}

	err = h.Services.OSBB.UpdateAnswer(userID, osbbID, pollID, &input)
	if err != nil {
		logger.Error("failed to update answer", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to update answer: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) deleteUserAnswer(c *gin.Context) {
	logger := h.Logger.Named("deletePollAnswer").WithContext(c)

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

	pollID, err := strconv.ParseUint(c.Param("pollID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse osbb id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to parse osbb id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
		return
	}

	err = h.Services.OSBB.DeleteAnswer(userID, osbbID, pollID)
	if err != nil {
		logger.Error("failed to delete answer", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to delete poll: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) getPollsResults(c *gin.Context) {
	logger := h.Logger.Named("getAllPollsAnswer").WithContext(c)

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

	pollID, err := strconv.ParseUint(c.Param("pollID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse osbb id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to parse poll id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
		return
	}

	polls, err := h.Services.OSBB.GetPollResult(userID, osbbID, pollID)
	if err != nil {
		logger.Error("failed to get polls result", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get polls result: %w", err))
		return
	}

	c.JSON(http.StatusOK, polls)
}

func (h *Handler) updateTestAnswer(c *gin.Context) {
	logger := h.Logger.Named("updateTestAnswer").WithContext(c)

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

	pollID, err := strconv.ParseUint(c.Param("pollID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse poll id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to parse poll id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
		return
	}

	testAnswerID, err := strconv.ParseUint(c.Param("testAnswerID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse test answer id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to parse test answer id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
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
	if err := h.validateJSONTags(body, entity.EventTestAnswerUpdatePayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventTestAnswerUpdatePayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to bind JSON: %w", err)).Code("Failed bind JSON").Kind(errs.Validation))
		return
	}

	err = h.Services.OSBB.UpdateTestAnswer(userID, osbbID, pollID, testAnswerID, input)
	if err != nil {
		logger.Error("failed to update test answer", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to update answer: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) deleteTestAnswer(c *gin.Context) {
	logger := h.Logger.Named("deleteTestAnswer").WithContext(c)

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

	pollID, err := strconv.ParseUint(c.Param("pollID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse poll id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to parse poll id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
		return
	}

	testAnswerID, err := strconv.ParseUint(c.Param("testAnswerID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse osbb id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to parse osbb id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
		return
	}

	err = h.Services.OSBB.DeleteTestAnswer(userID, osbbID, pollID, testAnswerID)
	if err != nil {
		logger.Error("failed to delete test answer", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to delete test answer: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
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
	if err := h.validateJSONTags(body, entity.EventPaymentPayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventPaymentPayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to bind JSON: %w", err)).Code("Failed bind JSON").Kind(errs.Validation))
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

func (h *Handler) getAllPayments(c *gin.Context) {
	logger := h.Logger.Named("getAllPayment").WithContext(c)

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

	payments, err := h.Services.OSBB.ListPayments(userID, osbbID)
	if err != nil {
		logger.Error("failed to get list payments", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get list payments: %w", err))
		return
	}

	c.JSON(http.StatusOK, payments)
}

func (h *Handler) updatePayment(c *gin.Context) {
	logger := h.Logger.Named("updatePayment").WithContext(c)

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

	paymentID, err := strconv.ParseUint(c.Param("paymentID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse payment id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to parse payment id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
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
	if err := h.validateJSONTags(body, entity.EventPaymentUpdatePayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventPaymentUpdatePayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to bind JSON: %w", err)).Code("Failed bind JSON").Kind(errs.Validation))
		return
	}

	err = h.Services.OSBB.UpdatePayment(userID, osbbID, paymentID, input)
	if err != nil {
		logger.Error("failed to add payment", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to add payment: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) deletePayment(c *gin.Context) {
	logger := h.Logger.Named("deletePayment").WithContext(c)

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

	paymentID, err := strconv.ParseUint(c.Param("paymentID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse payment id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to parse payment id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
		return
	}

	err = h.Services.OSBB.DeletePayment(userID, osbbID, paymentID)
	if err != nil {
		logger.Error("failed to delete payment", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to delete payment: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
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

	osbbID, err := strconv.ParseUint(c.Param("osbbID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse osbb id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to parse osbb id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
		return
	}

	paymentID, err := strconv.ParseUint(c.Param("paymentID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse payment id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to payment id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
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
	if err := h.validateJSONTags(body, entity.EventUserPurchaseUpdatePayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventUserPurchaseUpdatePayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to bind JSON: %w", err)).Code("Failed bind JSON").Kind(errs.Validation))
		return
	}

	err = h.Services.OSBB.UserUpdatePurchase(userID, osbbID, paymentID, input)
	if err != nil {
		logger.Error("failed to add user payment", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to add user payment: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) getAllPurchasesByUser(c *gin.Context) {
	logger := h.Logger.Named("getAllPurchasesByUser").WithContext(c)

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
	if err := h.validateJSONTags(body, entity.EventPurchaseFilterPayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventPurchaseFilterPayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to bind JSON: %w", err)).Code("Failed bind JSON").Kind(errs.Validation))
		return
	}

	purchases, err := h.Services.OSBB.ListPurchasesByUser(userID, osbbID, input)
	if err != nil {
		logger.Error("failed to get list purchases", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get list purchases: %w", err))
		return
	}

	c.JSON(http.StatusOK, purchases)
}

func (h *Handler) getAllPurchasesByOSBBHead(c *gin.Context) {
	logger := h.Logger.Named("getAllPurchasesByOSBBHead").WithContext(c)

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
	if err := h.validateJSONTags(body, entity.EventPurchaseOSBBHeadFilterPayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventPurchaseOSBBHeadFilterPayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to bind JSON: %w", err)).Code("Failed bind JSON").Kind(errs.Validation))
		return
	}

	purchases, err := h.Services.OSBB.ListPurchasesByOSBBHead(userID, osbbID, input)
	if err != nil {
		logger.Error("failed to get list purchases", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to get list purchases: %w", err))
		return
	}

	c.JSON(http.StatusOK, purchases)
}

func (h *Handler) updatePurchase(c *gin.Context) {
	logger := h.Logger.Named("updatePurchase").WithContext(c)

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

	paymentID, err := strconv.ParseUint(c.Param("paymentID"), 10, 64)
	if err != nil {
		logger.Error("failed to parse payment id", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to payment id: %w", err)).Code("Failed parse param").Kind(errs.Validation))
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
	if err := h.validateJSONTags(body, entity.EventUserPurchaseUpdatePayload{}); err != nil {
		logger.Error("failed to validate JSON tags", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to validate JSON tags: %w", err))
		return
	}

	var input entity.EventUserPurchaseUpdatePayload
	if err := c.BindJSON(&input); err != nil {
		logger.Error("failed to bind JSON", "error", err)
		h.sendErrResponse(c, h.Logger,
			errs.Err(fmt.Errorf("failed to bind JSON: %w", err)).Code("Failed bind JSON").Kind(errs.Validation))
		return
	}

	err = h.Services.OSBB.UpdatePurchase(userID, osbbID, paymentID, input)
	if err != nil {
		logger.Error("failed to add update user purchase", "error", err)
		h.sendErrResponse(c, h.Logger, fmt.Errorf("failed to update user purchase: %w", err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}
