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

type Shape2D struct {
  Resource
}

func (me *Shape2D) BaseClass() string {
  return "Shape2D"
}

func NewShape2D() *Shape2D {
  str := StringNameFromStr("Shape2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Shape2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Shape2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Shape2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Shape2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Shape2D) SetCustomSolverBias(bias float64, )  {
  classNameV := StringNameFromStr("Shape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_solver_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Shape2D) GetCustomSolverBias() float64 {
  classNameV := StringNameFromStr("Shape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_solver_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Shape2D) Collide(local_xform Transform2D, with_shape Shape2D, shape_xform Transform2D, ) bool {
  classNameV := StringNameFromStr("Shape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("collide")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3709843132) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{local_xform.AsCTypePtr(), with_shape.AsCTypePtr(), shape_xform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Shape2D) CollideWithMotion(local_xform Transform2D, local_motion Vector2, with_shape Shape2D, shape_xform Transform2D, shape_motion Vector2, ) bool {
  classNameV := StringNameFromStr("Shape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("collide_with_motion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2869556801) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{local_xform.AsCTypePtr(), local_motion.AsCTypePtr(), with_shape.AsCTypePtr(), shape_xform.AsCTypePtr(), shape_motion.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Shape2D) CollideAndGetContacts(local_xform Transform2D, with_shape Shape2D, shape_xform Transform2D, ) PackedVector2Array {
  classNameV := StringNameFromStr("Shape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("collide_and_get_contacts")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3056932662) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{local_xform.AsCTypePtr(), with_shape.AsCTypePtr(), shape_xform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Shape2D) CollideWithMotionAndGetContacts(local_xform Transform2D, local_motion Vector2, with_shape Shape2D, shape_xform Transform2D, shape_motion Vector2, ) PackedVector2Array {
  classNameV := StringNameFromStr("Shape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("collide_with_motion_and_get_contacts")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3620351573) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{local_xform.AsCTypePtr(), local_motion.AsCTypePtr(), with_shape.AsCTypePtr(), shape_xform.AsCTypePtr(), shape_motion.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Shape2D) Draw(canvas_item RID, color Color, )  {
  classNameV := StringNameFromStr("Shape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2948539648) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{canvas_item.AsCTypePtr(), color.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Shape2D) GetRect() Rect2 {
  classNameV := StringNameFromStr("Shape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1639390495) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
