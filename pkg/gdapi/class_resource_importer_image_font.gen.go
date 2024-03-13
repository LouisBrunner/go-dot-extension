// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ResourceImporterImageFont struct {
  obj gdc.ObjectPtr
}

func (me *ResourceImporterImageFont) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ResourceImporterImageFont) BaseClass() string {
  return "ResourceImporterImageFont"
}



// Enums

func (me *ResourceImporterImageFont) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterImageFont) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterImageFont) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
