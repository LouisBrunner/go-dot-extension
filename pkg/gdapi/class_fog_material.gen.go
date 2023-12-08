// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type FogMaterial struct {
  obj gdc.ObjectPtr
}

func (me *FogMaterial) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *FogMaterial) BaseClass() string {
  return "FogMaterial"
}
