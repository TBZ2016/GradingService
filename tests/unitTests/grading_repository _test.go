package uniTest

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"kawa/gradingservice/internal/app/grading/model"
	"kawa/gradingservice/internal/app/grading/repository"
)

var (
	testDatabaseName   = "test_database"
	testCollectionName = "grades"
)

const (
	testGradeID      = "testGradeID"
	testStudentID    = "testStudentID"
	testTeacherID    = "testTeacherID"
	testAssignmentID = "testAssignmentID"
	testCourseID     = "testCourseID"
	testClassID      = "testClassID"
)

var testGrade = model.Grade{
	GradeID:      testGradeID,
	StudentID:    testStudentID,
	TeacherID:    testTeacherID,
	AssignmentID: testAssignmentID,
	CreatedAt:    time.Now(),
	CourseID:     testCourseID,
	Grade:        85,
	IsPass:       true,
	ClassID:      testClassID,
}

func setupTestRepository(t *testing.T) (*repository.GradingRepository, func()) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	assert.NoError(t, err)

	db := client.Database(testDatabaseName)

	repo := repository.NewGradingRepository(db)

	// Clean up any existing test data
	_, err = repo.GetByCursusID(testCourseID)
	if err == nil {
		err = repo.DeleteById(testGradeID)
		assert.NoError(t, err)
	}

	// Return a cleanup function
	cleanup := func() {

		err := repo.DeleteById(testGradeID)
		assert.NoError(t, err)
		client.Disconnect(context.TODO())
	}

	return repo, cleanup
}

func TestGradingRepository_Create(t *testing.T) {
	repo, cleanup := setupTestRepository(t)
	defer cleanup()

	// Test Create
	err := repo.Create(&testGrade)
	assert.NoError(t, err)

	// Verify the grade is inserted
	insertedGrade, err := repo.GetById(testGradeID)
	assert.NoError(t, err)
	assert.NotNil(t, insertedGrade)
	assert.Equal(t, testGrade, *insertedGrade)
}

func TestGradingRepository_GetByStudentID(t *testing.T) {
	repo, cleanup := setupTestRepository(t)
	defer cleanup()

	// Insert test data
	err := repo.Create(&testGrade)
	assert.NoError(t, err)

	// Test GetByStudentID
	grades, err := repo.GetByStudentID(testStudentID)
	assert.NoError(t, err)
	assert.NotNil(t, grades)
	assert.Equal(t, 1, len(grades))
	assert.Equal(t, testGrade, grades[0])
}

func TestGradingRepository_GetByClass(t *testing.T) {
	repo, cleanup := setupTestRepository(t)
	defer cleanup()

	// Insert test data
	err := repo.Create(&testGrade)
	assert.NoError(t, err)

	// Test GetByClass
	grades, err := repo.GetByClass(testClassID)
	assert.NoError(t, err)
	assert.NotNil(t, grades)
	assert.Equal(t, 1, len(grades))
	assert.Equal(t, testGrade, grades[0])
}

func TestGradingRepository_GetById(t *testing.T) {
	repo, cleanup := setupTestRepository(t)
	defer cleanup()

	// Insert test data
	err := repo.Create(&testGrade)
	assert.NoError(t, err)

	// Test GetById
	foundGrade, err := repo.GetById(testGradeID)
	assert.NoError(t, err)
	assert.NotNil(t, foundGrade)

	// Compare without sub-second precision in CreatedAt
	expectedCreatedAt := testGrade.CreatedAt.Truncate(time.Second)
	actualCreatedAt := foundGrade.CreatedAt.Truncate(time.Second)

	assert.Equal(t, testGrade, *foundGrade, "Mismatch in grade details")
	assert.Equal(t, expectedCreatedAt, actualCreatedAt, "Mismatch in CreatedAt")
}

func TestGradingRepository_Update(t *testing.T) {
	repo, cleanup := setupTestRepository(t)
	defer cleanup()

	// Insert test data
	err := repo.Create(&testGrade)
	assert.NoError(t, err)

	// Update the test data
	testGrade.Grade = 90
	err = repo.Update(&testGrade)
	assert.NoError(t, err)

	// Verify the grade is updated
	updatedGrade, err := repo.GetById(testGradeID)
	assert.NoError(t, err)
	assert.NotNil(t, updatedGrade)
	assert.Equal(t, 90, updatedGrade.Grade)
}

func TestGradingRepository_DeleteById(t *testing.T) {
	repo, cleanup := setupTestRepository(t)
	defer cleanup()

	// Insert test data
	err := repo.Create(&testGrade)
	assert.NoError(t, err)

	// Test DeleteById
	err = repo.DeleteById(testGradeID)
	assert.NoError(t, err)

	// Verify the grade is deleted
	deletedGrade, err := repo.GetById(testGradeID)
	assert.NoError(t, err)
	assert.Nil(t, deletedGrade)
}
