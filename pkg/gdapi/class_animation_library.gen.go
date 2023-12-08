// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimationLibrary struct {
  obj gdc.ObjectPtr
}

func createAnimationLibrary(obj gdc.ObjectPtr) *AnimationLibrary {
  return &AnimationLibrary{
    obj: obj,
  }
}

func (me *AnimationLibrary) BaseClass() string {
  return "AnimationLibrary"
}
