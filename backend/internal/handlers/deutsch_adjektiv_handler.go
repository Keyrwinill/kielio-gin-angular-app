package handlers

import (
	"kielio-gin-angular-app/backend/internal/dto"
	"kielio-gin-angular-app/backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeutschAdjektivHandler struct {
	Service services.DeutschAdjektivService
}

func (h DeutschAdjektivHandler) GetDeutschAdjektiv(c *gin.Context) {

	selected_type := c.Query("type")

	result, err := h.Service.GetByType(selected_type)

	if err != nil {
		c.JSON(http.StatusInternalServerError,
			dto.Response[any]{
				Success:   false,
				ErrorCode: "9999",
				Message:   err.Error(),
				Data:      nil,
			})
		return
	}

	c.JSON(http.StatusOK,
		dto.Response[[]dto.DeutschAdjektivResponse]{
			Success:   true,
			ErrorCode: "0000",
			Message:   "",
			Data:      result,
		})
}
