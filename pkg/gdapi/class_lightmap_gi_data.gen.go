// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type LightmapGIData struct {
  obj gdc.ObjectPtr
}

func createLightmapGIData(obj gdc.ObjectPtr) *LightmapGIData {
  return &LightmapGIData{
    obj: obj,
  }
}

func (me *LightmapGIData) BaseClass() string {
  return "LightmapGIData"
}
