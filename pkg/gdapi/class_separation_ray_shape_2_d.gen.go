// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SeparationRayShape2D struct {
  obj gdc.ObjectPtr
}

func (me *SeparationRayShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SeparationRayShape2D) BaseClass() string {
  return "SeparationRayShape2D"
}
