// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BaseMaterial3D struct {
  obj gdc.ObjectPtr
}

func createBaseMaterial3D(obj gdc.ObjectPtr) *BaseMaterial3D {
  return &BaseMaterial3D{
    obj: obj,
  }
}

func (me *BaseMaterial3D) BaseClass() string {
  return "BaseMaterial3D"
}
