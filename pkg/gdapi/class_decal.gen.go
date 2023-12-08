// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Decal struct {
  obj gdc.ObjectPtr
}

func createDecal(obj gdc.ObjectPtr) *Decal {
  return &Decal{
    obj: obj,
  }
}

func (me *Decal) BaseClass() string {
  return "Decal"
}
