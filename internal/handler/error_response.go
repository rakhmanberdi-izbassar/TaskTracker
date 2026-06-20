package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rakhmanberdi-izbassar/TaskTracker/internal/apperrors"
)

func respondWithError(c *gin.Context, err error) {
	var appErr *apperrors.AppError

	if errors.As(err, &appErr) {
		c.JSON(appErr.Code, gin.H{"error": appErr.Message})
		return
	}

	c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
}
