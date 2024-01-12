package uniTest

// import (
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/google/uuid"
// 	"github.com/stretchr/testify/assert"

// 	"kawa/gradingservice/internal/app/grading/model"
// 	"kawa/gradingservice/internal/app/grading/usecase"
// 	"kawa/gradingservice/internal/app/grading/usecase/mocks"
// )

// func TestGradingHandler_GetGradesByCursusID(t *testing.T) {
// 	// Setup
// 	mockUseCase := &mocks.GradingUseCaseInterface{}
// 	handler := NewGradingHandler(mockUseCase)
// 	router := gin.Default()
// 	router.GET("/grades/cursus/:cursusId", handler.GetGradesByCursusID)

// 	// Test
// 	cursusID := uuid.New().String()
// 	mockUseCase.On("GetGradesByCursusID", cursusID).Return([]model.Grade{{GradeID: "1", CourseID: "math", Grade: 85}}, nil)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/grades/cursus/"+cursusID, nil)
// 	router.ServeHTTP(w, req)

// 	// Assert
// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response []model.Grade
// 	err := json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)

// 	assert.Len(t, response, 1)
// 	assert.Equal(t, "1", response[0].GradeID)
// 	assert.Equal(t, "math", response[0].CourseID)
// 	assert.Equal(t, 85, response[0].Grade)

// 	mockUseCase.AssertExpectations(t)
// }

// func TestGradingHandler_CreateGrade(t *testing.T) {
// 	// Setup
// 	mockUseCase := &mocks.GradingUseCaseInterface{}
// 	handler := NewGradingHandler(mockUseCase)
// 	router := gin.Default()
// 	router.POST("/grades", handler.CreateGrade)

// 	// Test
// 	grade := model.Grade{
// 		StudentID: "student123",
// 		CourseID:  "math",
// 		Grade:     90,
// 		IsPass:    true,
// 		ClassID:   "class456",
// 	}

// 	payload, err := json.Marshal(grade)
// 	assert.NoError(t, err)

// 	mockUseCase.On("CreateGrade", &grade).Return(nil)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("POST", "/grades", strings.NewReader(string(payload)))
// 	req.Header.Set("Content-Type", "application/json")
// 	router.ServeHTTP(w, req)

// 	// Assert
// 	assert.Equal(t, http.StatusCreated, w.Code)
// 	mockUseCase.AssertExpectations(t)
// }

// func TestGradingHandler_GetGradesByStudentID(t *testing.T) {
// 	// Setup
// 	mockUseCase := &mocks.GradingUseCaseInterface{}
// 	handler := NewGradingHandler(mockUseCase)
// 	router := gin.Default()
// 	router.GET("/grades/student/:studentId", handler.GetGradesByStudentID)

// 	// Test
// 	studentID := uuid.New().String()
// 	mockUseCase.On("GetGradesByStudentID", studentID).Return([]model.Grade{{GradeID: "2", CourseID: "history", Grade: 78}}, nil)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/grades/student/"+studentID, nil)
// 	router.ServeHTTP(w, req)

// 	// Assert
// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response []model.Grade
// 	err := json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)

// 	assert.Len(t, response, 1)
// 	assert.Equal(t, "2", response[0].GradeID)
// 	assert.Equal(t, "history", response[0].CourseID)
// 	assert.Equal(t, 78, response[0].Grade)

// 	mockUseCase.AssertExpectations(t)
// }

// func TestGradingHandler_GetGradesByClass(t *testing.T) {
// 	// Setup
// 	mockUseCase := &mocks.GradingUseCaseInterface{}
// 	handler := NewGradingHandler(mockUseCase)
// 	router := gin.Default()
// 	router.GET("/grades/class/:classId", handler.GetGradesByClass)

// 	// Test
// 	classID := uuid.New().String()
// 	mockUseCase.On("GetGradesByClass", classID).Return([]model.Grade{{GradeID: "3", CourseID: "science", Grade: 92}}, nil)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/grades/class/"+classID, nil)
// 	router.ServeHTTP(w, req)

// 	// Assert
// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response []model.Grade
// 	err := json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)

// 	assert.Len(t, response, 1)
// 	assert.Equal(t, "3", response[0].GradeID)
// 	assert.Equal(t, "science", response[0].CourseID)
// 	assert.Equal(t, 92, response[0].Grade)

// 	mockUseCase.AssertExpectations(t)
// }

// func TestGradingHandler_GetGradeByID(t *testing.T) {
// 	// Setup
// 	mockUseCase := &mocks.GradingUseCaseInterface{}
// 	handler := NewGradingHandler(mockUseCase)
// 	router := gin.Default()
// 	router.GET("/grades/:gradeId", handler.GetGradeByID)

// 	// Test
// 	gradeID := uuid.New().String()
// 	mockUseCase.On("GetGradeByID", gradeID).Return(&model.Grade{GradeID: "4", CourseID: "english", Grade: 88}, nil)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/grades/"+gradeID, nil)
// 	router.ServeHTTP(w, req)

// 	// Assert
// 	assert.Equal(t, http.StatusOK, w.Code)

// 	var response model.Grade
// 	err := json.Unmarshal(w.Body.Bytes(), &response)
// 	assert.NoError(t, err)

// 	assert.Equal(t, "4", response.GradeID)
// 	assert.Equal(t, "english", response.CourseID)
// 	assert.Equal(t, 88, response.Grade)

// 	mockUseCase.AssertExpectations(t)
// }

// func TestGradingHandler_UpdateGrade(t *testing.T) {
// 	// Setup
// 	mockUseCase := &mocks.GradingUseCaseInterface{}
// 	handler := NewGradingHandler(mockUseCase)
// 	router := gin.Default()
// 	router.PUT("/grades/:gradeId", handler.UpdateGrade)

// 	// Test
// 	gradeID := uuid.New().String()
// 	updatedGrade := model.Grade{GradeID: gradeID, Grade: 95}

// 	payload, err := json.Marshal(updatedGrade)
// 	assert.NoError(t, err)

// 	mockUseCase.On("UpdateGrade", &updatedGrade).Return(nil)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("PUT", "/grades/"+gradeID, strings.NewReader(string(payload)))
// 	req.Header.Set("Content-Type", "application/json")
// 	router.ServeHTTP(w, req)

// 	// Assert
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	mockUseCase.AssertExpectations(t)
// }

// func TestGradingHandler_DeleteGradeByID(t *testing.T) {
// 	// Setup
// 	mockUseCase := &mocks.GradingUseCaseInterface{}
// 	handler := NewGradingHandler(mockUseCase)
// 	router := gin.Default()
// 	router.DELETE("/grades/:gradeId", handler.DeleteGradeByID)

// 	// Test
// 	gradeID := uuid.New().String()
// 	mockUseCase.On("DeleteGradeByID", gradeID).Return(nil)

// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("DELETE", "/grades/"+gradeID, nil)
// 	router.ServeHTTP(w, req)

// 	// Assert
// 	assert.Equal(t, http.StatusNoContent, w.Code)
// 	mockUseCase.AssertExpectations(t)
// }
