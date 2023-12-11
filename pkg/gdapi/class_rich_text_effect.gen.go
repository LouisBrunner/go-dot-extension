// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RichTextEffect struct {
  obj gdc.ObjectPtr
}

func (me *RichTextEffect) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RichTextEffect) BaseClass() string {
  return "RichTextEffect"
}



// Enums

func (me *RichTextEffect) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RichTextEffect) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RichTextEffect) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
