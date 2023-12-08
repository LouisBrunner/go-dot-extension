// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GDScript struct {
  obj gdc.ObjectPtr
}

func createGDScript(obj gdc.ObjectPtr) *GDScript {
  return &GDScript{
    obj: obj,
  }
}

func (me *GDScript) BaseClass() string {
  return "GDScript"
}
