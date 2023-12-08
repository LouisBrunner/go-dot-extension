// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Label3D struct {
  obj gdc.ObjectPtr
}

func createLabel3D(obj gdc.ObjectPtr) *Label3D {
  return &Label3D{
    obj: obj,
  }
}

func (me *Label3D) BaseClass() string {
  return "Label3D"
}
