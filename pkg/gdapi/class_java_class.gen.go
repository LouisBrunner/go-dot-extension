// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type JavaClass struct {
  obj gdc.ObjectPtr
}

func createJavaClass(obj gdc.ObjectPtr) *JavaClass {
  return &JavaClass{
    obj: obj,
  }
}

func (me *JavaClass) BaseClass() string {
  return "JavaClass"
}
