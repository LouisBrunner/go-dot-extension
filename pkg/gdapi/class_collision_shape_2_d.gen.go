// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type CollisionShape2D struct {
  obj gdc.ObjectPtr
}

func (me *CollisionShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *CollisionShape2D) BaseClass() string {
  return "CollisionShape2D"
}



// Enums

func (me *CollisionShape2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *CollisionShape2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *CollisionShape2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *CollisionShape2D) SetShape(shape Shape2D, )  {
  classNameV := StringNameFromStr("CollisionShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 771364740) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shape.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionShape2D) GetShape() Shape2D {
  classNameV := StringNameFromStr("CollisionShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 522005891) // FIXME: should cache?
  var ret Shape2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionShape2D) SetDisabled(disabled bool, )  {
  classNameV := StringNameFromStr("CollisionShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionShape2D) IsDisabled() bool {
  classNameV := StringNameFromStr("CollisionShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_disabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionShape2D) SetOneWayCollision(enabled bool, )  {
  classNameV := StringNameFromStr("CollisionShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_one_way_collision")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionShape2D) IsOneWayCollisionEnabled() bool {
  classNameV := StringNameFromStr("CollisionShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_one_way_collision_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionShape2D) SetOneWayCollisionMargin(margin float32, )  {
  classNameV := StringNameFromStr("CollisionShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_one_way_collision_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionShape2D) GetOneWayCollisionMargin() float32 {
  classNameV := StringNameFromStr("CollisionShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_one_way_collision_margin")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *CollisionShape2D) SetDebugColor(color Color, )  {
  classNameV := StringNameFromStr("CollisionShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_debug_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2920490490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *CollisionShape2D) GetDebugColor() Color {
  classNameV := StringNameFromStr("CollisionShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_debug_color")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3444240500) // FIXME: should cache?
  var ret Color
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
