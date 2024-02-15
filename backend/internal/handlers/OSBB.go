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
	var input entity.InputOSBB

	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	// Check if there are any additional fields in the JSON body
	if err := h.validateJSONTags(body, entity.InputOSBB{}); err != nil {
		return
	}

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

	osbbID, err := strconv.ParseUint(c.Param("osbbID"), 10, 64)
	if err != nil {
		return
	}

	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	// Check if there are any additional fields in the JSON body
	if err := h.validateJSONTags(body, entity.InputAnnouncement{}); err != nil {
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

func (h *Handler) addPoll(c *gin.Context) {

	//userID, errs := getUserId(c)
	//if errs != nil {
	//	return
	//}

	osbbID, err := strconv.ParseUint(c.Param("osbbID"), 10, 64)
	if err != nil {
		return
	}

	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	// Check if there are any additional fields in the JSON body
	if err := h.validateJSONTags(body, entity.InputPoll{}); err != nil {
		fmt.Println(err)
		return
	}

	var input entity.InputPoll
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		return
	}

	poll, err := h.Services.OSBB.AddPoll(1, osbbID, input)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": poll.ID,
	})
}

func (h *Handler) addPollTest(c *gin.Context) {

	//userID, errs := getUserId(c)
	//if errs != nil {
	//	return
	//}

	osbbID, err := strconv.ParseUint(c.Param("osbbID"), 10, 64)
	if err != nil {
		return
	}

	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	// Check if there are any additional fields in the JSON body
	if err := h.validateJSONTags(body, entity.InputPollTest{}); err != nil {
		fmt.Println(err)
		return
	}

	var input entity.InputPollTest
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		return
	}

	poll, err := h.Services.OSBB.AddPollTest(1, osbbID, input)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": poll.ID,
	})
}

func (h *Handler) addPollAnswer(c *gin.Context) {

	//userID, errs := getUserId(c)
	//if errs != nil {
	//	return
	//}

	pollID, err := strconv.ParseUint(c.Param("pollID"), 10, 64)
	if err != nil {
		return
	}

	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	// Check if there are any additional fields in the JSON body
	if err := h.validateJSONTags(body, entity.InputPollAnswer{}); err != nil {
		fmt.Println(err)
		return
	}

	var input entity.InputPollAnswer
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		return
	}

	answer, err := h.Services.OSBB.AddPollAnswer(1, pollID, input)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": answer.ID,
	})
}

func (h *Handler) addPollAnswerTest(c *gin.Context) {

	//userID, errs := getUserId(c)
	//if errs != nil {
	//	return
	//}

	pollID, err := strconv.ParseUint(c.Param("pollID"), 10, 64)
	if err != nil {
		return
	}

	body, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
	// Check if there are any additional fields in the JSON body
	if err := h.validateJSONTags(body, entity.InputPollAnswerTest{}); err != nil {
		fmt.Println(err)
		return
	}

	var input entity.InputPollAnswerTest
	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		return
	}

	poll, err := h.Services.OSBB.AddPollAnswerTest(1, pollID, input)
	if err != nil {
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": poll.ID,
	})
}
