package usecase

import "kawa/gradingservice/internal/app/grading/model"

type GradingUseCaseInterface interface {
	GetByCursusID(cursusID int) ([]model.Grade, error)

	CreateGrade(gradeDTO *model.GradeDTO) error

	GetByStudentID(studentID int) ([]model.Grade, error)

	GetByClass(classID int) ([]model.Grade, error)

	GetGradeByID(gradeID int) (*model.Grade, error)

	UpdateGrade(gradeDTO *model.GradeDTO) error

	DeleteGradeByID(gradeID int) error
}
