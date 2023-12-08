// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Sprite2D struct {
  obj gdc.ObjectPtr
}

func (me *Sprite2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Sprite2D) BaseClass() string {
  return "Sprite2D"
}
