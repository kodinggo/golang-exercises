package subject

import "golang-second-project/model"

type english struct {
	subjectName string
}

// GetName returns name attributes
func (e english) GetSubjectName() string {
	return e.subjectName
}

func NewEnglish(subjectName string) model.Subject {
	return english{subjectName: subjectName}
}
