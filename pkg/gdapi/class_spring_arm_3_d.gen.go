// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SpringArm3D struct {
  obj gdc.ObjectPtr
}

func createSpringArm3D(obj gdc.ObjectPtr) *SpringArm3D {
  return &SpringArm3D{
    obj: obj,
  }
}

func (me *SpringArm3D) BaseClass() string {
  return "SpringArm3D"
}
