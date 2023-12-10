// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ResourceFormatSaver struct {
  obj gdc.ObjectPtr
}

func (me *ResourceFormatSaver) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ResourceFormatSaver) BaseClass() string {
  return "ResourceFormatSaver"
}



// Enums

func (me *ResourceFormatSaver) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceFormatSaver) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceFormatSaver) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties