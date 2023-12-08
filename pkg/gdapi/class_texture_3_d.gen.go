// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Texture3D struct {
  obj gdc.ObjectPtr
}

func (me *Texture3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Texture3D) BaseClass() string {
  return "Texture3D"
}
