// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForPlaceholderMaterialList struct {
}

var ptrsForPlaceholderMaterial ptrsForPlaceholderMaterialList

func initPlaceholderMaterialPtrs(iface gdc.Interface) {

  className := StringNameFromStr("PlaceholderMaterial")
  defer className.Destroy()
}

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
