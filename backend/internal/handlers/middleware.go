package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	UserCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	token, err := checkHeaderToken(c)
	if err != nil {
		return
	}
	isRevoked, err := h.checkIsRevoked(token)
	if err != nil {
		return
	}
	if isRevoked {
		return
	}

	id, err := h.Services.Token.ParseToken(token)
	if err != nil {
		return
	}
	c.Set(UserCtx, id)
}

// Function to validate JSON tags against the structure
func (h *Handler) validateJSONTags(body []byte, input interface{}) error {

	// Parse the JSON body into a map
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal(body, &jsonMap)
	if err != nil {
		return err
	}
	structType := reflect.TypeOf(input)

	// Iterate through the fields of the struct
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		// Get the JSON tag value for the field
		tagValue := field.Tag.Get("json")
		if _, ok := jsonMap["id"]; ok {
			return fmt.Errorf("invalid JSON tags: id")
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

		return fmt.Errorf("invalid JSON tags: %s", strings.Join(errorTags[:], ", "))
	}
	return nil
}

func getUserId(c *gin.Context) (uint64, error) {
	id, ok := c.Get(UserCtx)
	if !ok {
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(uint64)
	if !ok {
		return 0, errors.New("user id not found")
	}

	return idInt, nil
}

func checkHeaderToken(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}
	return headerParts[1], nil
}

func (h *Handler) checkIsRevoked(tokenValue string) (bool, error) {
	token, err := h.Services.Token.GetByToken(tokenValue)
	if err != nil {
		return true, err
	}
	if token.Revoked {
		return true, nil
	} else {
		return false, nil
	}
}