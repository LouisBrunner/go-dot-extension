// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimatableBody3D struct {
  obj gdc.ObjectPtr
}

func createAnimatableBody3D(obj gdc.ObjectPtr) *AnimatableBody3D {
  return &AnimatableBody3D{
    obj: obj,
  }
}

func (me *AnimatableBody3D) BaseClass() string {
  return "AnimatableBody3D"
}
