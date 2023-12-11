// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type OccluderPolygon2D struct {
  obj gdc.ObjectPtr
}

func (me *OccluderPolygon2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *OccluderPolygon2D) BaseClass() string {
  return "OccluderPolygon2D"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&closed), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OccluderPolygon2D) IsClosed() bool {
  classNameV := StringNameFromStr("OccluderPolygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_closed")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OccluderPolygon2D) SetCullMode(cull_mode OccluderPolygon2DCullMode, )  {
  classNameV := StringNameFromStr("OccluderPolygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cull_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3500863002) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&cull_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OccluderPolygon2D) GetCullMode() OccluderPolygon2DCullMode {
  classNameV := StringNameFromStr("OccluderPolygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cull_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 33931036) // FIXME: should cache?
  var ret OccluderPolygon2DCullMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *OccluderPolygon2D) SetPolygon(polygon PackedVector2Array, )  {
  classNameV := StringNameFromStr("OccluderPolygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(polygon.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *OccluderPolygon2D) GetPolygon() PackedVector2Array {
  classNameV := StringNameFromStr("OccluderPolygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961356807) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
