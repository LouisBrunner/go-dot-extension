// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Shape2D struct {
  obj gdc.ObjectPtr
}

func (me *Shape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Shape2D) BaseClass() string {
  return "Shape2D"
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

func  (me *Shape2D) SetCustomSolverBias(bias float32, )  {
  classNameV := StringNameFromStr("Shape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_custom_solver_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bias), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Shape2D) GetCustomSolverBias() float32 {
  classNameV := StringNameFromStr("Shape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_custom_solver_bias")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Shape2D) Collide(local_xform Transform2D, with_shape Shape2D, shape_xform Transform2D, ) bool {
  classNameV := StringNameFromStr("Shape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("collide")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3709843132) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(local_xform.AsCTypePtr()), gdc.ConstTypePtr(with_shape.AsCTypePtr()), gdc.ConstTypePtr(shape_xform.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Shape2D) CollideWithMotion(local_xform Transform2D, local_motion Vector2, with_shape Shape2D, shape_xform Transform2D, shape_motion Vector2, ) bool {
  classNameV := StringNameFromStr("Shape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("collide_with_motion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2869556801) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(local_xform.AsCTypePtr()), gdc.ConstTypePtr(local_motion.AsCTypePtr()), gdc.ConstTypePtr(with_shape.AsCTypePtr()), gdc.ConstTypePtr(shape_xform.AsCTypePtr()), gdc.ConstTypePtr(shape_motion.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Shape2D) CollideAndGetContacts(local_xform Transform2D, with_shape Shape2D, shape_xform Transform2D, ) PackedVector2Array {
  classNameV := StringNameFromStr("Shape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("collide_and_get_contacts")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3056932662) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(local_xform.AsCTypePtr()), gdc.ConstTypePtr(with_shape.AsCTypePtr()), gdc.ConstTypePtr(shape_xform.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Shape2D) CollideWithMotionAndGetContacts(local_xform Transform2D, local_motion Vector2, with_shape Shape2D, shape_xform Transform2D, shape_motion Vector2, ) PackedVector2Array {
  classNameV := StringNameFromStr("Shape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("collide_with_motion_and_get_contacts")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3620351573) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(local_xform.AsCTypePtr()), gdc.ConstTypePtr(local_motion.AsCTypePtr()), gdc.ConstTypePtr(with_shape.AsCTypePtr()), gdc.ConstTypePtr(shape_xform.AsCTypePtr()), gdc.ConstTypePtr(shape_motion.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Shape2D) Draw(canvas_item RID, color Color, )  {
  classNameV := StringNameFromStr("Shape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("draw")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2948539648) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(canvas_item.AsCTypePtr()), gdc.ConstTypePtr(color.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Shape2D) GetRect() Rect2 {
  classNameV := StringNameFromStr("Shape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rect")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1639390495) // FIXME: should cache?
  var ret Rect2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
