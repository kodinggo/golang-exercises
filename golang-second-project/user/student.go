package user

import (
	"fmt"
	"golang-second-project/model"
)

type student struct {
	name string

	english model.Subject
	math    model.Subject
}

// GetName returns name attributes
func (s student) GetName() string {
	return s.name
}

// SetName sets attributes name
func (s *student) SetName(name string) {
	s.name = name
}

// GetInfo returns information of student
func (s student) GetInfo() string {
	return fmt.Sprintf("%s menyukai subject %s\n", s.name, s.math.GetSubjectName())
}

func NewStudent(english, math model.Subject) model.Student {
	return &student{
		english: english,
		math:    math,
	}
}
