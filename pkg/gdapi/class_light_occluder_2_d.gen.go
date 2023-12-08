// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type LightOccluder2D struct {
  obj gdc.ObjectPtr
}

func (me *LightOccluder2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *LightOccluder2D) BaseClass() string {
  return "LightOccluder2D"
}
