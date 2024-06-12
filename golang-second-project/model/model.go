package model

// Student is interface of object student
type Student interface {
	GetName() string
	SetName(name string)
	GetInfo() string
}

// Teacher is interface of object teacher
type Teacher interface {
	GetName() string
	SetName(name string)
}

// Subject is interface of object subject
type Subject interface {
	GetSubjectName() string
}
