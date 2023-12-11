// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type LightmapperRD struct {
  obj gdc.ObjectPtr
}

func (me *LightmapperRD) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *LightmapperRD) BaseClass() string {
  return "LightmapperRD"
}



// Enums

func (me *LightmapperRD) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *LightmapperRD) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *LightmapperRD) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
