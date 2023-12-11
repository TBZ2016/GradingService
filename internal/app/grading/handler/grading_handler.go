package handler

import (
	"net/http"
	"strconv"

	"grading-service/internal/app/grading/usecase"

	"github.com/gin-gonic/gin"
)

type GradingHandler struct {
	GradingUseCase usecase.GradingUseCase
}

func NewGradingHandler(gradingUseCase usecase.GradingUseCase) *GradingHandler {
	return &GradingHandler{
		GradingUseCase: gradingUseCase,
	}
}

func (h *GradingHandler) GetByCursusID(c *gin.Context) {
	cursusID, err := strconv.Atoi(c.Param("cursusId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cursusId"})
		return
	}

	grades, err := h.GradingUseCase.GetGradesByCursusID(cursusID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, grades)
}
