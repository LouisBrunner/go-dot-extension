// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type PlaceholderMaterial struct {
  Material
}

func (me *PlaceholderMaterial) BaseClass() string {
  return "PlaceholderMaterial"
}

func NewPlaceholderMaterial() *PlaceholderMaterial {
  str := StringNameFromStr("PlaceholderMaterial") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PlaceholderMaterial{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *PlaceholderMaterial) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PlaceholderMaterial) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PlaceholderMaterial) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
