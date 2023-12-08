// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PointLight2D struct {
  obj gdc.ObjectPtr
}

func (me *PointLight2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PointLight2D) BaseClass() string {
  return "PointLight2D"
}
