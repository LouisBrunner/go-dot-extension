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
