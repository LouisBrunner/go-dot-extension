// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type TriangleMesh struct {
  RefCounted
}

func (me *TriangleMesh) BaseClass() string {
  return "TriangleMesh"
}

func NewTriangleMesh() *TriangleMesh {
  str := StringNameFromStr("TriangleMesh") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &TriangleMesh{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *TriangleMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *TriangleMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *TriangleMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
