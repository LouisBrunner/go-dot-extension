// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TextServerAdvanced struct {
  TextServerExtension
}

func (me *TextServerAdvanced) BaseClass() string {
  return "TextServerAdvanced"
}

func NewTextServerAdvanced() *TextServerAdvanced {
  str := StringNameFromStr("TextServerAdvanced") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &TextServerAdvanced{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *TextServerAdvanced) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TextServerAdvanced) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TextServerAdvanced) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
