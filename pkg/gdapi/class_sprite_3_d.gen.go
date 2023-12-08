// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Sprite3D struct {
  obj gdc.ObjectPtr
}

func (me *Sprite3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Sprite3D) BaseClass() string {
  return "Sprite3D"
}
