// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ResourceImporterShaderFile struct {
  obj gdc.ObjectPtr
}

func (me *ResourceImporterShaderFile) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ResourceImporterShaderFile) BaseClass() string {
  return "ResourceImporterShaderFile"
}



// Enums

func (me *ResourceImporterShaderFile) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterShaderFile) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterShaderFile) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
