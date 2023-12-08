// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PathFollow3D struct {
  obj gdc.ObjectPtr
}

func (me *PathFollow3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PathFollow3D) BaseClass() string {
  return "PathFollow3D"
}
