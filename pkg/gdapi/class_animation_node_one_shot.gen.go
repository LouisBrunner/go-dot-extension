// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeOneShot struct {
  obj gdc.ObjectPtr
}

func createAnimationNodeOneShot(obj gdc.ObjectPtr) *AnimationNodeOneShot {
  return &AnimationNodeOneShot{
    obj: obj,
  }
}

func (me *AnimationNodeOneShot) BaseClass() string {
  return "AnimationNodeOneShot"
}
