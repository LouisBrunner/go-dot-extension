// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Lightmapper struct {
  obj gdc.ObjectPtr
}

func createLightmapper(obj gdc.ObjectPtr) *Lightmapper {
  return &Lightmapper{
    obj: obj,
  }
}

func (me *Lightmapper) BaseClass() string {
  return "Lightmapper"
}
