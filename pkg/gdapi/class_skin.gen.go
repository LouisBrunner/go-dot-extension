// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Skin struct {
  obj gdc.ObjectPtr
}

func (me *Skin) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Skin) BaseClass() string {
  return "Skin"
}



// Enums

func (me *Skin) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Skin) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Skin) SetBindCount(bind_count int, )  {
  panic("TODO: implement")
}

func  (me *Skin) GetBindCount()  {
  panic("TODO: implement")
}

func  (me *Skin) AddBind(bone int, pose Transform3D, )  {
  panic("TODO: implement")
}

func  (me *Skin) AddNamedBind(name String, pose Transform3D, )  {
  panic("TODO: implement")
}

func  (me *Skin) SetBindPose(bind_index int, pose Transform3D, )  {
  panic("TODO: implement")
}

func  (me *Skin) GetBindPose(bind_index int, )  {
  panic("TODO: implement")
}

func  (me *Skin) SetBindName(bind_index int, name StringName, )  {
  panic("TODO: implement")
}

func  (me *Skin) GetBindName(bind_index int, )  {
  panic("TODO: implement")
}

func  (me *Skin) SetBindBone(bind_index int, bone int, )  {
  panic("TODO: implement")
}

func  (me *Skin) GetBindBone(bind_index int, )  {
  panic("TODO: implement")
}

func  (me *Skin) ClearBinds()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
