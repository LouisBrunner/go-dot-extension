// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CSGCombiner3D struct {
  CSGShape3D
}

func (me *CSGCombiner3D) BaseClass() string {
  return "CSGCombiner3D"
}

func NewCSGCombiner3D() *CSGCombiner3D {
  str := StringNameFromStr("CSGCombiner3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CSGCombiner3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *CSGCombiner3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CSGCombiner3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CSGCombiner3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
