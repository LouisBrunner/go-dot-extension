// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ResourceImporterOBJ struct {
  obj gdc.ObjectPtr
}

func (me *ResourceImporterOBJ) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ResourceImporterOBJ) BaseClass() string {
  return "ResourceImporterOBJ"
}



// Enums

func (me *ResourceImporterOBJ) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterOBJ) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterOBJ) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
