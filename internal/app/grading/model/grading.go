package model

type Grade struct {
	GradeID      int    `json:"gradeId"`
	StudentID    int    `json:"studentId"`
	TeacherID    int    `json:"teacherId"`
	AssignmentID int    `json:"assignmentId"`
	CreatedAt    string `json:"createdAt"`
	CourseID     int    `json:"courseId"`
	Grade        int    `json:"grade"`
	IsPass       bool   `json:"isPass"`
}
