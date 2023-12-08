// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type World3D struct {
  obj gdc.ObjectPtr
}

func (me *World3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *World3D) BaseClass() string {
  return "World3D"
}
