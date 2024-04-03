// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Lightmapper struct {
  RefCounted
}

func (me *Lightmapper) BaseClass() string {
  return "Lightmapper"
}

func NewLightmapper() *Lightmapper {
  str := StringNameFromStr("Lightmapper") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Lightmapper{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Lightmapper) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Lightmapper) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Lightmapper) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
