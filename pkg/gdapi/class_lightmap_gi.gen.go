// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type LightmapGI struct {
  obj gdc.ObjectPtr
}

func createLightmapGI(obj gdc.ObjectPtr) *LightmapGI {
  return &LightmapGI{
    obj: obj,
  }
}

func (me *LightmapGI) BaseClass() string {
  return "LightmapGI"
}
