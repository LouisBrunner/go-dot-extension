// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type StaticBody2D struct {
  obj gdc.ObjectPtr
}

func (me *StaticBody2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *StaticBody2D) BaseClass() string {
  return "StaticBody2D"
}
