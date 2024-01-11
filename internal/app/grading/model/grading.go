package model

import "time"

type Grade struct {
	GradeID      string    `json:"gradeId"`
	StudentID    string    `json:"studentId"`
	TeacherID    string    `json:"teacherId"`
	AssignmentID string    `json:"assignmentId"`
	CreatedAt    time.Time `json:"createdAt"`
	CourseID     string    `json:"courseId"`
	Grade        int       `json:"grade"`
	IsPass       bool      `json:"isPass"`
	ClassID      string    `json:"classId"`
}
