// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimatableBody2D struct {
  obj gdc.ObjectPtr
}

func (me *AnimatableBody2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimatableBody2D) BaseClass() string {
  return "AnimatableBody2D"
}
