package subject

import "golang-second-project/model"

type math struct {
	subjectName string
}

// GetName returns name attributes
func (m math) GetSubjectName() string {
	return m.subjectName
}

func NewMath(subjectName string) model.Subject {
	return math{subjectName: subjectName}
}
