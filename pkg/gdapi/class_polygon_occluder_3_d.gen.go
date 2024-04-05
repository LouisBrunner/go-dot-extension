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

type ptrsForPolygonOccluder3DList struct {
  fnSetPolygon gdc.MethodBindPtr
  fnGetPolygon gdc.MethodBindPtr
}

var ptrsForPolygonOccluder3D ptrsForPolygonOccluder3DList

func initPolygonOccluder3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("PolygonOccluder3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_polygon")
    defer methodName.Destroy()
    ptrsForPolygonOccluder3D.fnSetPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1509147220))
  }
  {
    methodName := StringNameFromStr("get_polygon")
    defer methodName.Destroy()
    ptrsForPolygonOccluder3D.fnGetPolygon = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2961356807))
  }
}

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
  cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygonOccluder3D.fnSetPolygon), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PolygonOccluder3D) GetPolygon() PackedVector2Array {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygonOccluder3D.fnGetPolygon), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
