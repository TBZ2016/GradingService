package repository

import "grading-service/internal/app/grading/model"

type GradingRepository interface {
	// GetByCursusID retrieves all grades by cursus ID.
	GetByCursusID(cursusID int) ([]model.Grade, error)

	// Create inserts a new grade into the repository.
	Create(grade *model.Grade) error

	// GetByStudentID retrieves grades for a specific student.
	GetByStudentID(studentID int) ([]model.Grade, error)

	// GetByClass retrieves grades for a specific class.
	GetByClass(classID int) ([]model.Grade, error)

	// GetById retrieves a grade by its ID.
	GetById(gradeID int) (*model.Grade, error)

	// Update updates an existing grade in the repository.
	Update(grade *model.Grade) error

	// DeleteById deletes a grade by its ID.
	DeleteById(gradeID int) error
}
