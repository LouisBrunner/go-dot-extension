// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SpriteBase3D struct {
  obj gdc.ObjectPtr
}

func (me *SpriteBase3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SpriteBase3D) BaseClass() string {
  return "SpriteBase3D"
}
