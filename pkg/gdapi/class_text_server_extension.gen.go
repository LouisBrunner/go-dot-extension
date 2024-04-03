// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TextServerExtension struct {
  TextServer
}

func (me *TextServerExtension) BaseClass() string {
  return "TextServerExtension"
}

func NewTextServerExtension() *TextServerExtension {
  str := StringNameFromStr("TextServerExtension") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &TextServerExtension{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *TextServerExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextServerExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextServerExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
