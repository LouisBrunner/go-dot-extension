// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SoftBody3D struct {
  obj gdc.ObjectPtr
}

func createSoftBody3D(obj gdc.ObjectPtr) *SoftBody3D {
  return &SoftBody3D{
    obj: obj,
  }
}

func (me *SoftBody3D) BaseClass() string {
  return "SoftBody3D"
}
