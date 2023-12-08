// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PackedDataContainer struct {
  obj gdc.ObjectPtr
}

func createPackedDataContainer(obj gdc.ObjectPtr) *PackedDataContainer {
  return &PackedDataContainer{
    obj: obj,
  }
}

func (me *PackedDataContainer) BaseClass() string {
  return "PackedDataContainer"
}
