// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Camera2D struct {
  obj gdc.ObjectPtr
}

func (me *Camera2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Camera2D) BaseClass() string {
  return "Camera2D"
}
