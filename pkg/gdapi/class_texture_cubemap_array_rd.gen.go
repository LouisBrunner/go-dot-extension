// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TextureCubemapArrayRD struct {
  obj gdc.ObjectPtr
}

func (me *TextureCubemapArrayRD) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *TextureCubemapArrayRD) BaseClass() string {
  return "TextureCubemapArrayRD"
}



// Enums

func (me *TextureCubemapArrayRD) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextureCubemapArrayRD) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextureCubemapArrayRD) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
