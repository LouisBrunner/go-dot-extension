// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Camera3D struct {
  obj gdc.ObjectPtr
}

func (me *Camera3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Camera3D) BaseClass() string {
  return "Camera3D"
}
