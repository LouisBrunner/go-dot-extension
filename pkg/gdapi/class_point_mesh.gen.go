// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PointMesh struct {
  PrimitiveMesh
}

func (me *PointMesh) BaseClass() string {
  return "PointMesh"
}



// Enums

func (me *PointMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PointMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PointMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
