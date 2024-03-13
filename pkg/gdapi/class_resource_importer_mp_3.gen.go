// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ResourceImporterMP3 struct {
  obj gdc.ObjectPtr
}

func (me *ResourceImporterMP3) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ResourceImporterMP3) BaseClass() string {
  return "ResourceImporterMP3"
}



// Enums

func (me *ResourceImporterMP3) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterMP3) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterMP3) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
