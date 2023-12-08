// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimatedTexture struct {
  obj gdc.ObjectPtr
}

func createAnimatedTexture(obj gdc.ObjectPtr) *AnimatedTexture {
  return &AnimatedTexture{
    obj: obj,
  }
}

func (me *AnimatedTexture) BaseClass() string {
  return "AnimatedTexture"
}
