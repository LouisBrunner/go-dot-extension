// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Environment struct {
  obj gdc.ObjectPtr
}

func createEnvironment(obj gdc.ObjectPtr) *Environment {
  return &Environment{
    obj: obj,
  }
}

func (me *Environment) BaseClass() string {
  return "Environment"
}
