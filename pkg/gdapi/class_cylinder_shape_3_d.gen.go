// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type CylinderShape3D struct {
  obj gdc.ObjectPtr
}

func createCylinderShape3D(obj gdc.ObjectPtr) *CylinderShape3D {
  return &CylinderShape3D{
    obj: obj,
  }
}

func (me *CylinderShape3D) BaseClass() string {
  return "CylinderShape3D"
}
