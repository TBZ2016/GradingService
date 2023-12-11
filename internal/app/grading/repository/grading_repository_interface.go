package repository

import "kawa/gradingservice/internal/app/grading/model"

type GradingRepositoryInterface interface {
	GetByCursusID(cursusID int) ([]model.Grade, error)

	Create(grade *model.Grade) error

	GetByStudentID(studentID int) ([]model.Grade, error)

	GetByClass(classID int) ([]model.Grade, error)

	GetById(gradeID int) (*model.Grade, error)

	Update(grade *model.Grade) error

	DeleteById(gradeID int) error
}
