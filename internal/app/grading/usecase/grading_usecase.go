package usecase

import (
	"kawa/gradingservice/internal/app/grading/model"
	"kawa/gradingservice/internal/app/grading/repository"
)

type GradingUseCase struct {
	gradingRepo repository.GradingRepositoryInterface
}

func NewGradingUseCase(repo repository.GradingRepositoryInterface) *GradingUseCase {
	return &GradingUseCase{
		gradingRepo: repo,
	}
}

func (uc *GradingUseCase) GetGradesByCursusID(cursusID int) ([]model.Grade, error) {
	return uc.gradingRepo.GetByCursusID(cursusID)
}

func (uc *GradingUseCase) CreateGrade(grade *model.Grade) error {
	return uc.gradingRepo.Create(grade)
}

func (uc *GradingUseCase) GetGradesByStudentID(studentID int) ([]model.Grade, error) {
	return uc.gradingRepo.GetByStudentID(studentID)
}

func (uc *GradingUseCase) GetGradesByClass(classID int) ([]model.Grade, error) {
	return uc.gradingRepo.GetByClass(classID)
}

func (uc *GradingUseCase) GetGradeByID(gradeID int) (*model.Grade, error) {
	return uc.gradingRepo.GetById(gradeID)
}

func (uc *GradingUseCase) UpdateGrade(grade *model.Grade) error {
	return uc.gradingRepo.Update(grade)
}

func (uc *GradingUseCase) DeleteGradeByID(gradeID int) error {
	return uc.gradingRepo.DeleteById(gradeID)
}
