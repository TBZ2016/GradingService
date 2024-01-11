package usecase

import "kawa/gradingservice/internal/app/grading/model"

type GradingUseCaseInterface interface {
	GetGradesByCursusID(cursusID string) ([]model.Grade, error)
	CreateGrade(grade *model.Grade) error
	GetGradesByStudentID(studentID string) ([]model.Grade, error)
	GetGradesByClass(classID string) ([]model.Grade, error)
	GetGradeByID(gradeID string) (*model.Grade, error)
	UpdateGrade(grade *model.Grade) error
	DeleteGradeByID(gradeID string) error
}
