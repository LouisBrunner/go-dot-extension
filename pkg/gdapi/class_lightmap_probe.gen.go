// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type LightmapProbe struct {
  obj gdc.ObjectPtr
}

func createLightmapProbe(obj gdc.ObjectPtr) *LightmapProbe {
  return &LightmapProbe{
    obj: obj,
  }
}

func (me *LightmapProbe) BaseClass() string {
  return "LightmapProbe"
}
