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

type ptrsForCollisionShape2DList struct {
  fnSetShape gdc.MethodBindPtr
  fnGetShape gdc.MethodBindPtr
  fnSetDisabled gdc.MethodBindPtr
  fnIsDisabled gdc.MethodBindPtr
  fnSetOneWayCollision gdc.MethodBindPtr
  fnIsOneWayCollisionEnabled gdc.MethodBindPtr
  fnSetOneWayCollisionMargin gdc.MethodBindPtr
  fnGetOneWayCollisionMargin gdc.MethodBindPtr
  fnSetDebugColor gdc.MethodBindPtr
  fnGetDebugColor gdc.MethodBindPtr
}

var ptrsForCollisionShape2D ptrsForCollisionShape2DList

func initCollisionShape2DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("CollisionShape2D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_shape")
    defer methodName.Destroy()
    ptrsForCollisionShape2D.fnSetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 771364740))
  }
  {
    methodName := StringNameFromStr("get_shape")
    defer methodName.Destroy()
    ptrsForCollisionShape2D.fnGetShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 522005891))
  }
  {
    methodName := StringNameFromStr("set_disabled")
    defer methodName.Destroy()
    ptrsForCollisionShape2D.fnSetDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_disabled")
    defer methodName.Destroy()
    ptrsForCollisionShape2D.fnIsDisabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_one_way_collision")
    defer methodName.Destroy()
    ptrsForCollisionShape2D.fnSetOneWayCollision = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("is_one_way_collision_enabled")
    defer methodName.Destroy()
    ptrsForCollisionShape2D.fnIsOneWayCollisionEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_one_way_collision_margin")
    defer methodName.Destroy()
    ptrsForCollisionShape2D.fnSetOneWayCollisionMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_one_way_collision_margin")
    defer methodName.Destroy()
    ptrsForCollisionShape2D.fnGetOneWayCollisionMargin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_debug_color")
    defer methodName.Destroy()
    ptrsForCollisionShape2D.fnSetDebugColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2920490490))
  }
  {
    methodName := StringNameFromStr("get_debug_color")
    defer methodName.Destroy()
    ptrsForCollisionShape2D.fnGetDebugColor = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3444240500))
  }
}

type CollisionShape2D struct {
  Node2D
}

func (me *CollisionShape2D) BaseClass() string {
  return "CollisionShape2D"
}

func NewCollisionShape2D() *CollisionShape2D {
  str := StringNameFromStr("CollisionShape2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &CollisionShape2D{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{shape.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionShape2D.fnSetShape), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionShape2D) GetShape() Shape2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewShape2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionShape2D.fnGetShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *CollisionShape2D) SetDisabled(disabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&disabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionShape2D.fnSetDisabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionShape2D) IsDisabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionShape2D.fnIsDisabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CollisionShape2D) SetOneWayCollision(enabled bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionShape2D.fnSetOneWayCollision), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionShape2D) IsOneWayCollisionEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionShape2D.fnIsOneWayCollisionEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CollisionShape2D) SetOneWayCollisionMargin(margin float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&margin) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionShape2D.fnSetOneWayCollisionMargin), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionShape2D) GetOneWayCollisionMargin() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionShape2D.fnGetOneWayCollisionMargin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *CollisionShape2D) SetDebugColor(color Color, )  {
  cargs := []gdc.ConstTypePtr{color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionShape2D.fnSetDebugColor), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *CollisionShape2D) GetDebugColor() Color {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewColor()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForCollisionShape2D.fnGetDebugColor), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
