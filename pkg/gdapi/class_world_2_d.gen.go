// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type World2D struct {
  obj gdc.ObjectPtr
}

func (me *World2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *World2D) BaseClass() string {
  return "World2D"
}
