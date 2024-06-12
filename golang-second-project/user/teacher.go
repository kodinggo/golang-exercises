package user

import (
	"golang-second-project/model"
)

type teacher struct {
	name string
}

// GetName returns name attributes
func (t teacher) GetName() string {
	return t.name
}

// SetName sets attributes name
func (s *teacher) SetName(name string) {
	s.name = name
}

func NewTeacher() model.Teacher {
	return &teacher{}
}
