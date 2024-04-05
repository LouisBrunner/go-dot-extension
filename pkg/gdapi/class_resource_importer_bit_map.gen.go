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

type ptrsForResourceImporterBitMapList struct {
}

var ptrsForResourceImporterBitMap ptrsForResourceImporterBitMapList

func initResourceImporterBitMapPtrs(iface gdc.Interface) {

  className := StringNameFromStr("ResourceImporterBitMap")
  defer className.Destroy()
}

type ResourceImporterBitMap struct {
  ResourceImporter
}

func (me *ResourceImporterBitMap) BaseClass() string {
  return "ResourceImporterBitMap"
}

func NewResourceImporterBitMap() *ResourceImporterBitMap {
  str := StringNameFromStr("ResourceImporterBitMap") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ResourceImporterBitMap{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ResourceImporterBitMap) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceImporterBitMap) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceImporterBitMap) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
