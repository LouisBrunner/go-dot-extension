// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationPlayer struct {
  obj gdc.ObjectPtr
}

func (me *AnimationPlayer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationPlayer) BaseClass() string {
  return "AnimationPlayer"
}
