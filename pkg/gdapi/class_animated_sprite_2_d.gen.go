// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimatedSprite2D struct {
  obj gdc.ObjectPtr
}

func createAnimatedSprite2D(obj gdc.ObjectPtr) *AnimatedSprite2D {
  return &AnimatedSprite2D{
    obj: obj,
  }
}

func (me *AnimatedSprite2D) BaseClass() string {
  return "AnimatedSprite2D"
}
