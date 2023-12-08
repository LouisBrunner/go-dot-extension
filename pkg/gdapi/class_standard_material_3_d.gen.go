// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StandardMaterial3D struct {
  obj gdc.ObjectPtr
}

func createStandardMaterial3D(obj gdc.ObjectPtr) *StandardMaterial3D {
  return &StandardMaterial3D{
    obj: obj,
  }
}

func (me *StandardMaterial3D) BaseClass() string {
  return "StandardMaterial3D"
}
