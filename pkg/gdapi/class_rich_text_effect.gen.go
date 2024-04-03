// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RichTextEffect struct {
  Resource
}

func (me *RichTextEffect) BaseClass() string {
  return "RichTextEffect"
}

func NewRichTextEffect() *RichTextEffect {
  str := StringNameFromStr("RichTextEffect") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RichTextEffect{}
  obj.SetBaseObject(objPtr)
  return obj
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
