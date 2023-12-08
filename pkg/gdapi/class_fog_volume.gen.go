// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type FogVolume struct {
  obj gdc.ObjectPtr
}

func (me *FogVolume) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *FogVolume) BaseClass() string {
  return "FogVolume"
}
