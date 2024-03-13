// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ResourceImporterDynamicFont struct {
  obj gdc.ObjectPtr
}

func (me *ResourceImporterDynamicFont) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ResourceImporterDynamicFont) BaseClass() string {
  return "ResourceImporterDynamicFont"
}



// Enums

func (me *ResourceImporterDynamicFont) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterDynamicFont) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterDynamicFont) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
