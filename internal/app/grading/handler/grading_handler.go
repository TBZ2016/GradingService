package handler

import (
	"net/http"

	"kawa/gradingservice/internal/app/grading/model"
	"kawa/gradingservice/internal/app/grading/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	cursusIDStr := c.Param("cursusId")
	cursusID, err := uuid.Parse(cursusIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cursus ID"})
		return
	}

	grades, err := h.gradingUseCase.GetGradesByCursusID(cursusID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(grades) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No grades found for the specified student ID"})
		return
	}

	c.JSON(http.StatusOK, grades)
}

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

func (h *GradingHandler) GetGradesByStudentID(c *gin.Context) {
	studentIDStr := c.Param("studentId")
	studentID, err := uuid.Parse(studentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	grades, err := h.gradingUseCase.GetGradesByStudentID(studentID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(grades) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No grades found for the specified student ID"})
		return
	}

	c.JSON(http.StatusOK, grades)
}

func (h *GradingHandler) GetGradesByClass(c *gin.Context) {
	classIDStr := c.Param("classId")
	classID, err := uuid.Parse(classIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Class ID is not a GUID"})
		return
	}

	grades, err := h.gradingUseCase.GetGradesByClass(classID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(grades) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No grades found for the specified class ID"})
		return
	}

	c.JSON(http.StatusOK, grades)
}

func (h *GradingHandler) GetGradeByID(c *gin.Context) {
	gradeIDStr := c.Param("gradeId")
	gradeID, err := uuid.Parse(gradeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid grade ID"})
		return
	}

	grade, err := h.gradingUseCase.GetGradeByID(gradeID.String())
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

func (h *GradingHandler) UpdateGrade(c *gin.Context) {
	gradeID := c.Param("gradeId")

	var updatedGrade model.Grade
	if err := c.ShouldBindJSON(&updatedGrade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the gradeId from the URL parameter
	updatedGrade.GradeID = gradeID

	err := h.gradingUseCase.UpdateGrade(&updatedGrade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Grade updated"})
}

func (h *GradingHandler) DeleteGradeByID(c *gin.Context) {
	gradeIDStr := c.Param("gradeId")
	gradeID, err := uuid.Parse(gradeIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid grade ID"})
		return
	}

	err = h.gradingUseCase.DeleteGradeByID(gradeID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Grade deleted", "gradeId": gradeID.String()})
}
