// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CollisionPolygon2D struct {
  obj gdc.ObjectPtr
}

func (me *CollisionPolygon2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CollisionPolygon2D) BaseClass() string {
  return "CollisionPolygon2D"
}



// Enums

type CollisionPolygon2DBuildMode int
const (
  CollisionPolygon2DBuildModeBuildSolids CollisionPolygon2DBuildMode = 0
  CollisionPolygon2DBuildModeBuildSegments CollisionPolygon2DBuildMode = 1
)

func (me *CollisionPolygon2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CollisionPolygon2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CollisionPolygon2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CollisionPolygon2D) SetPolygon(polygon PackedVector2Array, )  {
  classNameV := StringNameFromStr("CollisionPolygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(polygon.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionPolygon2D) GetPolygon() PackedVector2Array {
  classNameV := StringNameFromStr("CollisionPolygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_polygon")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961356807) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionPolygon2D) SetBuildMode(build_mode CollisionPolygon2DBuildMode, )  {
  classNameV := StringNameFromStr("CollisionPolygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_build_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2780803135) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&build_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionPolygon2D) GetBuildMode() CollisionPolygon2DBuildMode {
  classNameV := StringNameFromStr("CollisionPolygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_build_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3044948800) // FIXME: should cache?
  var ret CollisionPolygon2DBuildMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionPolygon2D) SetDisabled(disabled bool, )  {
  classNameV := StringNameFromStr("CollisionPolygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionPolygon2D) IsDisabled() bool {
  classNameV := StringNameFromStr("CollisionPolygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionPolygon2D) SetOneWayCollision(enabled bool, )  {
  classNameV := StringNameFromStr("CollisionPolygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_one_way_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionPolygon2D) IsOneWayCollisionEnabled() bool {
  classNameV := StringNameFromStr("CollisionPolygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_one_way_collision_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionPolygon2D) SetOneWayCollisionMargin(margin float32, )  {
  classNameV := StringNameFromStr("CollisionPolygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_one_way_collision_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionPolygon2D) GetOneWayCollisionMargin() float32 {
  classNameV := StringNameFromStr("CollisionPolygon2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_one_way_collision_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
