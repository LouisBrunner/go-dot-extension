// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SpriteBase3D struct {
  obj gdc.ObjectPtr
}

func createSpriteBase3D(obj gdc.ObjectPtr) *SpriteBase3D {
  return &SpriteBase3D{
    obj: obj,
  }
}

func (me *SpriteBase3D) BaseClass() string {
  return "SpriteBase3D"
}
