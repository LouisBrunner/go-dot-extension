// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationNodeTimeSeek struct {
  obj gdc.ObjectPtr
}

func createAnimationNodeTimeSeek(obj gdc.ObjectPtr) *AnimationNodeTimeSeek {
  return &AnimationNodeTimeSeek{
    obj: obj,
  }
}

func (me *AnimationNodeTimeSeek) BaseClass() string {
  return "AnimationNodeTimeSeek"
}
