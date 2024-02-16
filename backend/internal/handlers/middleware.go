package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"residential-registration/backend/pkg/errs"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	UserCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	logger := h.Logger.Named("userIdentity").WithContext(c)

	token, err := h.checkHeaderToken(c)
	if err != nil {
		h.sendErrResponse(c, logger, fmt.Errorf("failed to check header token: %w", err))
		return
	}
	isRevoked, err := h.checkIsRevoked(token)
	if err != nil {
		h.sendErrResponse(c, logger, fmt.Errorf("failed to check is token revoked: %w", err))
		return
	}
	if isRevoked {
		h.sendErrResponse(c, logger, fmt.Errorf("failed to get access: %w",
			errs.M("token is revoked").Code("Failed to get access").Kind(errs.Private)))
		return
	}

	id, err := h.Services.Token.ParseToken(token)
	if err != nil {
		h.sendErrResponse(c, logger, fmt.Errorf("failed to parse token: %w", err))
		return
	}
	c.Set(UserCtx, id)
}

// Function to validate JSON tags against the structure
func (h *Handler) validateJSONTags(body []byte, input interface{}) error {
	logger := h.Logger.Named("validateJSONTags")
	// Parse the JSON body into a map
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(body, &jsonMap)
	if err != nil {
		logger.Error("failed validate JSOn tags", "error", errors.New("error to unmarshal"))
		return errs.Err(errors.New("error to unmarshal")).Code("Failed validate JSOn tags").Kind(errs.Private)
	}
	structType := reflect.TypeOf(input)

	// Iterate through the fields of the struct
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		// Get the JSON tag value for the field
		tagValue := field.Tag.Get("json")
		if _, ok := jsonMap["id"]; ok {
			logger.Error("failed validate JSOn tags", "error", errors.New("invalid JSON tags: id"))
			return errs.Err(errors.New("invalid JSON tags: id")).Code("Failed validate JSOn tags").Kind(errs.Private)

		}
		// Check if the JSON tag is not empty
		if tagValue != "" {
			// Check if the field exists in the struct
			if _, ok := jsonMap[tagValue]; ok {
				delete(jsonMap, tagValue)
			}
		}
	}
	if len(jsonMap) != 0 {
		var errorTags []string
		for key := range jsonMap {
			errorTags = append(errorTags, key)
		}
		logger.Error("failed validate JSOn tags", "error", fmt.Errorf("invalid JSON tags: %s", strings.Join(errorTags[:], ", ")))
		return errs.Err(fmt.Errorf("invalid JSON tags: %s", strings.Join(errorTags[:], ", "))).Code("Failed validate JSOn tags").Kind(errs.Validation)
	}
	return nil
}

func (h *Handler) getUserId(c *gin.Context) (uint64, error) {
	logger := h.Logger.Named("getUserId")
	id, ok := c.Get(UserCtx)
	if !ok {
		logger.Error("failed user authentication", "error", errors.New("user id not found"))
		return 0, errs.Err(errors.New("user id not found")).Code("Failed user authentication").Kind(errs.Validation)
	}
	idInt, ok := id.(uint64)
	if !ok {
		logger.Error("failed user authentication", "error", errors.New("error parse id"))
		return 0, errs.Err(errors.New("error parse id")).Code("Failed user authentication").Kind(errs.Validation)
	}

	return idInt, nil
}

func (h *Handler) checkHeaderToken(c *gin.Context) (string, error) {
	logger := h.Logger.Named("checkHeaderToken")
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		logger.Error("invalid token", "error", errors.New("empty auth header"))
		return "", errs.Err(errors.New("empty auth header")).Code("Invalid token").Kind(errs.Validation)
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		logger.Error("invalid token", "error", errors.New("invalid auth header"))
		return "", errs.Err(errors.New("invalid auth header")).Code("Invalid token").Kind(errs.Validation)
	}

	if len(headerParts[1]) == 0 {
		logger.Error("invalid token", "error", errors.New("token is empty"))
		return "", errs.Err(errors.New("token is empty")).Code("Invalid token").Kind(errs.Validation)
	}
	return headerParts[1], nil
}

func (h *Handler) checkIsRevoked(tokenValue string) (bool, error) {
	logger := h.Logger.Named("checkIsRevoked").
		With("token_value", tokenValue)
	token, err := h.Services.Token.GetByToken(tokenValue)
	if err != nil {
		logger.Error("failed to get token")
		return true, err
	}
	if token.Revoked {
		return true, nil
	} else {
		return false, nil
	}
}
