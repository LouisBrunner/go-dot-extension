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

type Node2D struct {
  CanvasItem
}

func (me *Node2D) BaseClass() string {
  return "Node2D"
}

func NewNode2D() *Node2D {
  str := StringNameFromStr("Node2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Node2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Node2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Node2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Node2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Node2D) SetPosition(position Vector2, )  {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node2D) SetRotation(radians float64, )  {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node2D) SetRotationDegrees(degrees float64, )  {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_rotation_degrees")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node2D) SetSkew(radians float64, )  {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_skew")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node2D) SetScale(scale Vector2, )  {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{scale.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node2D) GetPosition() Vector2 {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node2D) GetRotation() float64 {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Node2D) GetRotationDegrees() float64 {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rotation_degrees")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Node2D) GetSkew() float64 {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skew")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Node2D) GetScale() Vector2 {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node2D) Rotate(radians float64, )  {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("rotate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node2D) MoveLocalX(delta float64, scaled bool, )  {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_local_x")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2087892650) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delta) , gdc.ConstTypePtr(&scaled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node2D) MoveLocalY(delta float64, scaled bool, )  {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("move_local_y")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2087892650) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delta) , gdc.ConstTypePtr(&scaled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node2D) Translate(offset Vector2, )  {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("translate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node2D) GlobalTranslate(offset Vector2, )  {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("global_translate")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{offset.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node2D) ApplyScale(ratio Vector2, )  {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("apply_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{ratio.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node2D) SetGlobalPosition(position Vector2, )  {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_global_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node2D) GetGlobalPosition() Vector2 {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node2D) SetGlobalRotation(radians float64, )  {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_global_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node2D) SetGlobalRotationDegrees(degrees float64, )  {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_global_rotation_degrees")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&degrees) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node2D) GetGlobalRotation() float64 {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Node2D) GetGlobalRotationDegrees() float64 {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_rotation_degrees")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Node2D) SetGlobalSkew(radians float64, )  {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_global_skew")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&radians) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node2D) GetGlobalSkew() float64 {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_skew")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Node2D) SetGlobalScale(scale Vector2, )  {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_global_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{scale.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node2D) GetGlobalScale() Vector2 {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_global_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node2D) SetTransform(xform Transform2D, )  {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2761652528) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{xform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node2D) SetGlobalTransform(xform Transform2D, )  {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_global_transform")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2761652528) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{xform.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node2D) LookAt(point Vector2, )  {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("look_at")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Node2D) GetAngleTo(point Vector2, ) float64 {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_angle_to")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2276447920) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Node2D) ToLocal(global_point Vector2, ) Vector2 {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("to_local")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2656412154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{global_point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node2D) ToGlobal(local_point Vector2, ) Vector2 {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("to_global")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2656412154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{local_point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Node2D) GetRelativeTransformToParent(parent Node, ) Transform2D {
  classNameV := StringNameFromStr("Node2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_relative_transform_to_parent")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 904556875) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{parent.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
