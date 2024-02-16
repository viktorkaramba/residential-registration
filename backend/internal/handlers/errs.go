package handlers

import (
	"errors"
	"residential-registration/backend/pkg/errs"
	"residential-registration/backend/pkg/logging"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func (h *Handler) sendErrResponse(c *gin.Context, logger logging.Logger, err error) {

	var e *errs.Error
	ok := errors.As(err, &e)
	if !ok || e.IsServer() {
		logger.Error("unexpected error", "error", err)
	}
	c.AbortWithStatusJSON(e.HTTPStatusCode(), ErrorResponse{err.Error()})
}
