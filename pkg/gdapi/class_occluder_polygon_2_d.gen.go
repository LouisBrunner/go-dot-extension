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

type OccluderPolygon2D struct {
  Resource
}

func (me *OccluderPolygon2D) BaseClass() string {
  return "OccluderPolygon2D"
}

func NewOccluderPolygon2D() *OccluderPolygon2D {
  str := StringNameFromStr("OccluderPolygon2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &OccluderPolygon2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type OccluderPolygon2DCullMode int
const (
  OccluderPolygon2DCullModeCullDisabled OccluderPolygon2DCullMode = 0
  OccluderPolygon2DCullModeCullClockwise OccluderPolygon2DCullMode = 1
  OccluderPolygon2DCullModeCullCounterClockwise OccluderPolygon2DCullMode = 2
)

func (me *OccluderPolygon2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *OccluderPolygon2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *OccluderPolygon2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *OccluderPolygon2D) SetClosed(closed bool, )  {
  classNameV := StringNameFromStr("OccluderPolygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_closed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&closed) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OccluderPolygon2D) IsClosed() bool {
  classNameV := StringNameFromStr("OccluderPolygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_closed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *OccluderPolygon2D) SetCullMode(cull_mode OccluderPolygon2DCullMode, )  {
  classNameV := StringNameFromStr("OccluderPolygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cull_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3500863002) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cull_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OccluderPolygon2D) GetCullMode() OccluderPolygon2DCullMode {
  classNameV := StringNameFromStr("OccluderPolygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cull_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 33931036) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret OccluderPolygon2DCullMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *OccluderPolygon2D) SetPolygon(polygon PackedVector2Array, )  {
  classNameV := StringNameFromStr("OccluderPolygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{polygon.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *OccluderPolygon2D) GetPolygon() PackedVector2Array {
  classNameV := StringNameFromStr("OccluderPolygon2D")
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
