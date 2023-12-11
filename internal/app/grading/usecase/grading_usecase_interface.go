package usecase

import "kawa/gradingservice/internal/app/grading/model"

type GradingUseCaseInterface interface {
	GetGradesByCursusID(cursusID int) ([]model.Grade, error)
	CreateGrade(grade *model.Grade) error
	GetGradesByStudentID(studentID int) ([]model.Grade, error)
	GetGradesByClass(classID int) ([]model.Grade, error)
	GetGradeByID(gradeID int) (*model.Grade, error)
	UpdateGrade(grade *model.Grade) error
	DeleteGradeByID(gradeID int) error
}
