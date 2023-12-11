package handler

import (
	"net/http"
	"strconv"

	"kawa/gradingservice/internal/app/grading/model"
	"kawa/gradingservice/internal/app/grading/usecase"

	"github.com/gin-gonic/gin"
)

type GradingHandler struct {
	gradingUseCase usecase.GradingUseCaseInterface
}

func NewGradingHandler(useCase usecase.GradingUseCaseInterface) *GradingHandler {
	return &GradingHandler{
		gradingUseCase: useCase,
	}
}
func (h *GradingHandler) GetGradesByCursusID(c *gin.Context) {
	cursusID, err := strconv.Atoi(c.Param("cursusId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cursus ID"})
		return
	}

	grades, err := h.gradingUseCase.GetGradesByCursusID(cursusID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, grades)
}

// CreateGrade handles the creation of a new grade.
func (h *GradingHandler) CreateGrade(c *gin.Context) {
	var grade model.Grade
	if err := c.ShouldBindJSON(&grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.gradingUseCase.CreateGrade(&grade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Grade created"})
}

// GetGradesByStudentID handles the retrieval of grades by student ID.
func (h *GradingHandler) GetGradesByStudentID(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("studentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	grades, err := h.gradingUseCase.GetGradesByStudentID(studentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, grades)
}

// GetGradesByClass handles the retrieval of grades by class.
func (h *GradingHandler) GetGradesByClass(c *gin.Context) {
	classID, err := strconv.Atoi(c.Param("class"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid class ID"})
		return
	}

	grades, err := h.gradingUseCase.GetGradesByClass(classID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, grades)
}

// GetGradeByID handles the retrieval of a grade by ID.
func (h *GradingHandler) GetGradeByID(c *gin.Context) {
	gradeID, err := strconv.Atoi(c.Param("gradeId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid grade ID"})
		return
	}

	grade, err := h.gradingUseCase.GetGradeByID(gradeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if grade == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Grade not found"})
		return
	}

	c.JSON(http.StatusOK, grade)
}

// UpdateGrade handles the update of an existing grade.
func (h *GradingHandler) UpdateGrade(c *gin.Context) {
	var updatedGrade model.Grade
	if err := c.ShouldBindJSON(&updatedGrade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.gradingUseCase.UpdateGrade(&updatedGrade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Grade updated"})
}

// DeleteGradeByID handles the deletion of a grade by ID.
func (h *GradingHandler) DeleteGradeByID(c *gin.Context) {
	gradeID, err := strconv.Atoi(c.Param("gradeId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid grade ID"})
		return
	}

	err = h.gradingUseCase.DeleteGradeByID(gradeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Grade deleted"})
}
