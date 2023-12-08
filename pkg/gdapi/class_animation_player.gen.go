// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationPlayer struct {
  obj gdc.ObjectPtr
}

func createAnimationPlayer(obj gdc.ObjectPtr) *AnimationPlayer {
  return &AnimationPlayer{
    obj: obj,
  }
}

func (me *AnimationPlayer) BaseClass() string {
  return "AnimationPlayer"
}
