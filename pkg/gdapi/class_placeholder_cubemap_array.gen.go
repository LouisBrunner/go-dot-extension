// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PlaceholderCubemapArray struct {
  PlaceholderTextureLayered
}

func (me *PlaceholderCubemapArray) BaseClass() string {
  return "PlaceholderCubemapArray"
}

func NewPlaceholderCubemapArray() *PlaceholderCubemapArray {
  str := StringNameFromStr("PlaceholderCubemapArray") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PlaceholderCubemapArray{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *PlaceholderCubemapArray) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PlaceholderCubemapArray) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PlaceholderCubemapArray) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
