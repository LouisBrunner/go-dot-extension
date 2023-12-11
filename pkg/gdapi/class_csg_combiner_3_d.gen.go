// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CSGCombiner3D struct {
  obj gdc.ObjectPtr
}

func (me *CSGCombiner3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CSGCombiner3D) BaseClass() string {
  return "CSGCombiner3D"
}



// Enums

func (me *CSGCombiner3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CSGCombiner3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CSGCombiner3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
