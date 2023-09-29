package models

type Student struct {
	StudentId   int64  `json:"studentId,omitempty"`
	StudentName string `json:"student_name"`
	CollegeName string `json:"collegeName"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
}
