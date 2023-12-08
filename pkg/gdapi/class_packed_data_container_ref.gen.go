// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PackedDataContainerRef struct {
  obj gdc.ObjectPtr
}

func createPackedDataContainerRef(obj gdc.ObjectPtr) *PackedDataContainerRef {
  return &PackedDataContainerRef{
    obj: obj,
  }
}

func (me *PackedDataContainerRef) BaseClass() string {
  return "PackedDataContainerRef"
}
