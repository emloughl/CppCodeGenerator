package cppcomponents

// ClassImplementation ... Implements File
type ClassImplementation struct {
	inheritedInterface Interface
	name               string
}

func (c ClassImplementation) newClassImplementation() *ClassImplementation {

	return &c
}