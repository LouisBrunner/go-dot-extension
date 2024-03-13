// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ResourceImporterBMFont struct {
  obj gdc.ObjectPtr
}

func (me *ResourceImporterBMFont) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ResourceImporterBMFont) BaseClass() string {
  return "ResourceImporterBMFont"
}



// Enums

func (me *ResourceImporterBMFont) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterBMFont) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterBMFont) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
