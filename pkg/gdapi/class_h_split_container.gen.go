// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type HSplitContainer struct {
  obj gdc.ObjectPtr
}

func (me *HSplitContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *HSplitContainer) BaseClass() string {
  return "HSplitContainer"
}



// Enums

func (me *HSplitContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *HSplitContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *HSplitContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties