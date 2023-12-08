// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeSync struct {
  obj gdc.ObjectPtr
}

func createAnimationNodeSync(obj gdc.ObjectPtr) *AnimationNodeSync {
  return &AnimationNodeSync{
    obj: obj,
  }
}

func (me *AnimationNodeSync) BaseClass() string {
  return "AnimationNodeSync"
}
