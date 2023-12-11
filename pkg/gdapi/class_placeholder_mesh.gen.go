// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PlaceholderMesh struct {
  obj gdc.ObjectPtr
}

func (me *PlaceholderMesh) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PlaceholderMesh) BaseClass() string {
  return "PlaceholderMesh"
}



// Enums

func (me *PlaceholderMesh) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PlaceholderMesh) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PlaceholderMesh) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PlaceholderMesh) SetAabb(aabb AABB, )  {
  classNameV := StringNameFromStr("PlaceholderMesh")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_aabb")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 259215842) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(aabb.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
