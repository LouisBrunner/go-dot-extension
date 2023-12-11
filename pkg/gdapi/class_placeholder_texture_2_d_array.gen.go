// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PlaceholderTexture2DArray struct {
  obj gdc.ObjectPtr
}

func (me *PlaceholderTexture2DArray) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PlaceholderTexture2DArray) BaseClass() string {
  return "PlaceholderTexture2DArray"
}



// Enums

func (me *PlaceholderTexture2DArray) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PlaceholderTexture2DArray) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PlaceholderTexture2DArray) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
