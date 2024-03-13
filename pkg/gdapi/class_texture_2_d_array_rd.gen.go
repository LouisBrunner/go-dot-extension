// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Texture2DArrayRD struct {
  obj gdc.ObjectPtr
}

func (me *Texture2DArrayRD) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Texture2DArrayRD) BaseClass() string {
  return "Texture2DArrayRD"
}



// Enums

func (me *Texture2DArrayRD) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Texture2DArrayRD) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Texture2DArrayRD) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
