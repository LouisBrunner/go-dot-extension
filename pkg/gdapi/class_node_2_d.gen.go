// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Node2D struct {
  obj gdc.ObjectPtr
}

func (me *Node2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Node2D) BaseClass() string {
  return "Node2D"
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
  panic("TODO: implement")
}

func  (me *Node2D) SetRotation(radians float32, )  {
  panic("TODO: implement")
}

func  (me *Node2D) SetRotationDegrees(degrees float32, )  {
  panic("TODO: implement")
}

func  (me *Node2D) SetSkew(radians float32, )  {
  panic("TODO: implement")
}

func  (me *Node2D) SetScale(scale Vector2, )  {
  panic("TODO: implement")
}

func  (me *Node2D) GetPosition()  {
  panic("TODO: implement")
}

func  (me *Node2D) GetRotation()  {
  panic("TODO: implement")
}

func  (me *Node2D) GetRotationDegrees()  {
  panic("TODO: implement")
}

func  (me *Node2D) GetSkew()  {
  panic("TODO: implement")
}

func  (me *Node2D) GetScale()  {
  panic("TODO: implement")
}

func  (me *Node2D) Rotate(radians float32, )  {
  panic("TODO: implement")
}

func  (me *Node2D) MoveLocalX(delta float32, scaled bool, )  {
  panic("TODO: implement")
}

func  (me *Node2D) MoveLocalY(delta float32, scaled bool, )  {
  panic("TODO: implement")
}

func  (me *Node2D) Translate(offset Vector2, )  {
  panic("TODO: implement")
}

func  (me *Node2D) GlobalTranslate(offset Vector2, )  {
  panic("TODO: implement")
}

func  (me *Node2D) ApplyScale(ratio Vector2, )  {
  panic("TODO: implement")
}

func  (me *Node2D) SetGlobalPosition(position Vector2, )  {
  panic("TODO: implement")
}

func  (me *Node2D) GetGlobalPosition()  {
  panic("TODO: implement")
}

func  (me *Node2D) SetGlobalRotation(radians float32, )  {
  panic("TODO: implement")
}

func  (me *Node2D) SetGlobalRotationDegrees(degrees float32, )  {
  panic("TODO: implement")
}

func  (me *Node2D) GetGlobalRotation()  {
  panic("TODO: implement")
}

func  (me *Node2D) GetGlobalRotationDegrees()  {
  panic("TODO: implement")
}

func  (me *Node2D) SetGlobalSkew(radians float32, )  {
  panic("TODO: implement")
}

func  (me *Node2D) GetGlobalSkew()  {
  panic("TODO: implement")
}

func  (me *Node2D) SetGlobalScale(scale Vector2, )  {
  panic("TODO: implement")
}

func  (me *Node2D) GetGlobalScale()  {
  panic("TODO: implement")
}

func  (me *Node2D) SetTransform(xform Transform2D, )  {
  panic("TODO: implement")
}

func  (me *Node2D) SetGlobalTransform(xform Transform2D, )  {
  panic("TODO: implement")
}

func  (me *Node2D) LookAt(point Vector2, )  {
  panic("TODO: implement")
}

func  (me *Node2D) GetAngleTo(point Vector2, )  {
  panic("TODO: implement")
}

func  (me *Node2D) ToLocal(global_point Vector2, )  {
  panic("TODO: implement")
}

func  (me *Node2D) ToGlobal(local_point Vector2, )  {
  panic("TODO: implement")
}

func  (me *Node2D) GetRelativeTransformToParent(parent Node, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
