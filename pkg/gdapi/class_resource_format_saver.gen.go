// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ResourceFormatSaver struct {
  RefCounted
}

func (me *ResourceFormatSaver) BaseClass() string {
  return "ResourceFormatSaver"
}

func NewResourceFormatSaver() *ResourceFormatSaver {
  str := StringNameFromStr("ResourceFormatSaver") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ResourceFormatSaver{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ResourceFormatSaver) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ResourceFormatSaver) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ResourceFormatSaver) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
