// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Sprite2D struct {
  obj gdc.ObjectPtr
}

func createSprite2D(obj gdc.ObjectPtr) *Sprite2D {
  return &Sprite2D{
    obj: obj,
  }
}

func (me *Sprite2D) BaseClass() string {
  return "Sprite2D"
}
