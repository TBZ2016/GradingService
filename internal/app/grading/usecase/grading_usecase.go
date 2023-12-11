package usecase

import (
	"kawa/gradingservice/internal/app/grading/model"
	"kawa/gradingservice/internal/app/grading/repository"
)

type GradingUseCase struct {
	gradingRepository repository.GradingRepository
}

func NewGradingUseCase(gradingRepository repository.GradingRepository) *GradingUseCase {
	return &GradingUseCase{
		gradingRepository: gradingRepository,
	}
}

func (uc *GradingUseCase) GetGradesByCursusID(cursusID int) ([]model.Grade, error) {
	return uc.gradingRepository.GetByCursusID(cursusID)
}

func (uc *GradingUseCase) CreateGrade(grade *model.Grade) error {
	return uc.gradingRepository.Create(grade)
}

func (uc *GradingUseCase) GetGradesByStudentID(studentID int) ([]model.Grade, error) {
	return uc.gradingRepository.GetByStudentID(studentID)
}

func (uc *GradingUseCase) GetGradesByClass(classID int) ([]model.Grade, error) {
	return uc.gradingRepository.GetByClass(classID)
}

func (uc *GradingUseCase) GetGradeByID(gradeID int) (*model.Grade, error) {
	return uc.gradingRepository.GetById(gradeID)
}

func (uc *GradingUseCase) UpdateGrade(grade *model.Grade) error {
	return uc.gradingRepository.Update(grade)
}

func (uc *GradingUseCase) DeleteGradeByID(gradeID int) error {
	return uc.gradingRepository.DeleteById(gradeID)
}
