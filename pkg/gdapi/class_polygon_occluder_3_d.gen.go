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

type PolygonOccluder3D struct {
  Occluder3D
}

func (me *PolygonOccluder3D) BaseClass() string {
  return "PolygonOccluder3D"
}

func NewPolygonOccluder3D() *PolygonOccluder3D {
  str := StringNameFromStr("PolygonOccluder3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PolygonOccluder3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *PolygonOccluder3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PolygonOccluder3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PolygonOccluder3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PolygonOccluder3D) SetPolygon(polygon PackedVector2Array, )  {
  classNameV := StringNameFromStr("PolygonOccluder3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PolygonOccluder3D) GetPolygon() PackedVector2Array {
  classNameV := StringNameFromStr("PolygonOccluder3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961356807) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
