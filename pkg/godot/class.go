package godot

import "github.com/LouisBrunner/go-dot-extension/pkg/gdextension"

type class struct {
	name       string
	parentName string
}

func NewClass(name string, parentName string) gdextension.ClassConstructor {
	return func() gdextension.Class {
		return &class{
			name:       name,
			parentName: parentName,
		}
	}
}

func (me *class) ClassName() string {
	return me.name
}

func (me *class) ParentClassName() string {
	return me.parentName
}
