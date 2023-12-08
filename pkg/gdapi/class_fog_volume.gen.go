// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type FogVolume struct {
  obj gdc.ObjectPtr
}

func createFogVolume(obj gdc.ObjectPtr) *FogVolume {
  return &FogVolume{
    obj: obj,
  }
}

func (me *FogVolume) BaseClass() string {
  return "FogVolume"
}
