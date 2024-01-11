package uniTest

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"kawa/gradingservice/internal/app/grading/model"
	"kawa/gradingservice/internal/app/grading/usecase"
)

type mockGradingRepository struct {
	grades map[string]model.Grade
}

func (m *mockGradingRepository) GetByCursusID(cursusID string) ([]model.Grade, error) {
	var grades []model.Grade
	for _, grade := range m.grades {
		if grade.CourseID == cursusID {
			grades = append(grades, grade)
		}
	}
	return grades, nil
}

func (m *mockGradingRepository) GetByStudentID(studentID string) ([]model.Grade, error) {
	var grades []model.Grade
	for _, grade := range m.grades {
		if grade.StudentID == studentID {
			grades = append(grades, grade)
		}
	}
	return grades, nil
}

func (m *mockGradingRepository) GetByClass(classID string) ([]model.Grade, error) {
	var grades []model.Grade
	for _, grade := range m.grades {
		if grade.ClassID == classID {
			grades = append(grades, grade)
		}
	}
	return grades, nil
}

func (m *mockGradingRepository) GetById(gradeID string) (*model.Grade, error) {
	grade, exists := m.grades[gradeID]
	if !exists {
		return nil, nil
	}
	return &grade, nil
}

func (m *mockGradingRepository) Create(grade *model.Grade) error {
	m.grades[grade.GradeID] = *grade
	return nil
}

func (m *mockGradingRepository) Update(grade *model.Grade) error {
	m.grades[grade.GradeID] = *grade
	return nil
}

func (m *mockGradingRepository) DeleteById(gradeID string) error {
	delete(m.grades, gradeID)
	return nil
}

func TestGradingUseCase_GetGradesByCursusID(t *testing.T) {
	mockRepo := &mockGradingRepository{
		grades: map[string]model.Grade{
			"1": {GradeID: "1", CourseID: "math"},
			"2": {GradeID: "2", CourseID: "english"},
		},
	}

	useCase := usecase.NewGradingUseCase(mockRepo)

	cursusID := "math"
	grades, err := useCase.GetGradesByCursusID(cursusID)

	assert.NoError(t, err)
	assert.NotNil(t, grades)
	assert.Equal(t, 1, len(grades))
	assert.Equal(t, "math", grades[0].CourseID)
}

func TestGradingUseCase_GetGradesByStudentID(t *testing.T) {
	mockRepo := &mockGradingRepository{
		grades: map[string]model.Grade{
			"1": {GradeID: "1", StudentID: "student1"},
			"2": {GradeID: "2", StudentID: "student2"},
		},
	}

	useCase := usecase.NewGradingUseCase(mockRepo)

	studentID := "student1"
	grades, err := useCase.GetGradesByStudentID(studentID)

	assert.NoError(t, err)
	assert.NotNil(t, grades)
	assert.Equal(t, 1, len(grades))
	assert.Equal(t, "student1", grades[0].StudentID)
}

func TestGradingUseCase_GetGradesByClass(t *testing.T) {
	mockRepo := &mockGradingRepository{
		grades: map[string]model.Grade{
			"1": {GradeID: "1", ClassID: "classA"},
			"2": {GradeID: "2", ClassID: "classB"},
		},
	}

	useCase := usecase.NewGradingUseCase(mockRepo)

	classID := "classA"
	grades, err := useCase.GetGradesByClass(classID)

	assert.NoError(t, err)
	assert.NotNil(t, grades)
	assert.Equal(t, 1, len(grades))
	assert.Equal(t, "classA", grades[0].ClassID)
}

func TestGradingUseCase_GetGradeByID(t *testing.T) {
	mockRepo := &mockGradingRepository{
		grades: map[string]model.Grade{
			"1": {GradeID: "1", CourseID: "math"},
			"2": {GradeID: "2", CourseID: "english"},
		},
	}

	useCase := usecase.NewGradingUseCase(mockRepo)

	gradeID := "1"
	foundGrade, err := useCase.GetGradeByID(gradeID)

	assert.NoError(t, err)
	assert.NotNil(t, foundGrade)
	assert.Equal(t, "math", foundGrade.CourseID)
}

func TestGradingUseCase_CreateGrade(t *testing.T) {
	mockRepo := &mockGradingRepository{
		grades: make(map[string]model.Grade),
	}

	useCase := usecase.NewGradingUseCase(mockRepo)

	testGrade := &model.Grade{
		StudentID: "student1",
		TeacherID: "teacher1",
		CourseID:  "math",
		ClassID:   "classA",
	}

	err := useCase.CreateGrade(testGrade)

	assert.NoError(t, err)
	assert.NotEmpty(t, testGrade.GradeID)
	assert.NotNil(t, mockRepo.grades[testGrade.GradeID])
}

func TestGradingUseCase_UpdateGrade(t *testing.T) {
	mockRepo := &mockGradingRepository{
		grades: map[string]model.Grade{
			"1": {GradeID: "1", CourseID: "math", Grade: 85},
		},
	}

	useCase := usecase.NewGradingUseCase(mockRepo)

	testGrade := &model.Grade{
		GradeID:   "1",
		Grade:     90,
		IsPass:    true,
		CreatedAt: time.Now(),
	}

	err := useCase.UpdateGrade(testGrade)

	assert.NoError(t, err)
	assert.Equal(t, 90, mockRepo.grades["1"].Grade)
}

func TestGradingUseCase_UpdateGrade_NotFound(t *testing.T) {
	mockRepo := &mockGradingRepository{
		grades: make(map[string]model.Grade),
	}

	useCase := usecase.NewGradingUseCase(mockRepo)

	testGrade := &model.Grade{
		GradeID: "nonexistent",
		Grade:   90,
	}

	err := useCase.UpdateGrade(testGrade)

	assert.Error(t, err)
	assert.EqualError(t, err, "Grade with ID nonexistent not found")
}

func TestGradingUseCase_DeleteGradeByID(t *testing.T) {
	mockRepo := &mockGradingRepository{
		grades: map[string]model.Grade{
			"1": {GradeID: "1", CourseID: "math"},
		},
	}

	useCase := usecase.NewGradingUseCase(mockRepo)

	gradeID := "1"
	err := useCase.DeleteGradeByID(gradeID)

	assert.NoError(t, err)
	assert.NotContains(t, mockRepo.grades, gradeID)
}

func TestGradingUseCase_Validation_Error(t *testing.T) {
	mockRepo := &mockGradingRepository{
		grades: make(map[string]model.Grade),
	}

	useCase := usecase.NewGradingUseCase(mockRepo)

	testGrade := &model.Grade{
		// Missing required fields
	}

	err := useCase.CreateGrade(testGrade)

	assert.Error(t, err)
	assert.EqualError(t, err, "StudentID, TeacherID, ClassID, and CourseID are required fields")
}

func TestGradingUseCase_Validation_GradeOutOfRange(t *testing.T) {
	mockRepo := &mockGradingRepository{
		grades: make(map[string]model.Grade),
	}

	useCase := usecase.NewGradingUseCase(mockRepo)

	testGrade := &model.Grade{
		StudentID: "student1",
		TeacherID: "teacher1",
		CourseID:  "math",
		ClassID:   "classA",
		Grade:     150, // Grade out of range
	}

	err := useCase.CreateGrade(testGrade)

	assert.NoError(t, err)
	assert.Equal(t, 0, testGrade.Grade)
}
