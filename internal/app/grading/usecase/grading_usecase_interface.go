package usecase

import "grading-service/internal/app/grading/model"

type GradingUseCase interface {
	GetGradesByCursusID(cursusID int) ([]model.Grade, error)
	CreateGrade(grade *model.Grade) error
	// Add other use case methods as needed
}
