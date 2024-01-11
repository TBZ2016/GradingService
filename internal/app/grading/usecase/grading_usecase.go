package usecase

import (
	"fmt"
	"kawa/gradingservice/internal/app/grading/model"
	"kawa/gradingservice/internal/app/grading/repository"
	"time"

	"github.com/google/uuid"
)

type GradingUseCase struct {
	gradingRepo repository.GradingRepositoryInterface
}

func NewGradingUseCase(repo repository.GradingRepositoryInterface) *GradingUseCase {
	return &GradingUseCase{
		gradingRepo: repo,
	}
}

func (uc *GradingUseCase) GetGradesByCursusID(cursusID string) ([]model.Grade, error) {
	return uc.gradingRepo.GetByCursusID(cursusID)
}

func (uc *GradingUseCase) GetGradesByStudentID(studentID string) ([]model.Grade, error) {
	return uc.gradingRepo.GetByStudentID(studentID)
}

func (uc *GradingUseCase) GetGradesByClass(classID string) ([]model.Grade, error) {
	return uc.gradingRepo.GetByClass(classID)
}

func (uc *GradingUseCase) GetGradeByID(gradeID string) (*model.Grade, error) {
	return uc.gradingRepo.GetById(gradeID)
}

func (uc *GradingUseCase) CreateGrade(grade *model.Grade) error {
	if err := uc.validateGrade(grade, true); err != nil {
		return err
	}

	if grade.CreatedAt.IsZero() {
		grade.CreatedAt = time.Now()
	}

	return uc.gradingRepo.Create(grade)
}

func (uc *GradingUseCase) UpdateGrade(grade *model.Grade) error {
	existingGrade, err := uc.GetGradeByID(grade.GradeID)
	if err != nil {
		return err
	}
	if existingGrade == nil {
		return fmt.Errorf("Grade with ID %s not found", grade.GradeID)
	}

	if err := uc.validateGrade(grade, false); err != nil {
		return err
	}
	return uc.gradingRepo.Update(grade)
}

func (uc *GradingUseCase) DeleteGradeByID(gradeID string) error {
	existingGrade, err := uc.GetGradeByID(gradeID)
	if err != nil {
		return err
	}
	if existingGrade == nil {
		return fmt.Errorf("Grade with ID %s not found", gradeID)
	}
	return uc.gradingRepo.DeleteById(gradeID)
}

func (uc *GradingUseCase) validateGrade(grade *model.Grade, isCreate bool) error {

	// Check if required fields are filled
	if grade.StudentID == "" || grade.TeacherID == "" || grade.CourseID == "" || grade.ClassID == "" {
		return fmt.Errorf("StudentID, TeacherID, ClassID, and CourseID are required fields")
	}

	if isCreate {
		if grade.GradeID != "" {
			return fmt.Errorf("GradeID should not be provided for create operation")
		}

		if grade.Grade < 0 || grade.Grade > 100 {
			grade.Grade = 0
		}
		grade.GradeID = uuid.New().String()
	} else {

		if grade.GradeID == "" {
			return fmt.Errorf("GradeID is required for update operation")
		}

		// Check if Grade value is within the valid range
		if grade.Grade < 0 || grade.Grade > 100 {
			return fmt.Errorf("Grade value must be between 0 and 100")
		}
	}

	return nil
}
