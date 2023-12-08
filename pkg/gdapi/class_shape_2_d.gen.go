// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Shape2D struct {
  obj gdc.ObjectPtr
}

func (me *Shape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Shape2D) BaseClass() string {
  return "Shape2D"
}
