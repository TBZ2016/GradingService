package repository

import "kawa/gradingservice/internal/app/grading/model"

type GradingRepositoryInterface interface {
	GetByCursusID(cursusID string) ([]model.Grade, error)

	Create(grade *model.Grade) error

	GetByStudentID(studentID string) ([]model.Grade, error)

	GetByClass(classID string) ([]model.Grade, error)

	GetById(gradeID string) (*model.Grade, error)

	Update(grade *model.Grade) error

	DeleteById(gradeID string) error
}
