// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimatableBody2D struct {
  obj gdc.ObjectPtr
}

func createAnimatableBody2D(obj gdc.ObjectPtr) *AnimatableBody2D {
  return &AnimatableBody2D{
    obj: obj,
  }
}

func (me *AnimatableBody2D) BaseClass() string {
  return "AnimatableBody2D"
}
