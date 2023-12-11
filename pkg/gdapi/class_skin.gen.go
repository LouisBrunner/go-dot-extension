// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func (me *Skin) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Skin) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Skin) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Skin) SetBindCount(bind_count int, )  {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bind_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bind_count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Skin) GetBindCount() int {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bind_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Skin) AddBind(bone int, pose Transform3D, )  {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_bind")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3616898986) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone), gdc.ConstTypePtr(pose.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Skin) AddNamedBind(name String, pose Transform3D, )  {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_named_bind")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3154712474) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), gdc.ConstTypePtr(pose.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Skin) SetBindPose(bind_index int, pose Transform3D, )  {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bind_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3616898986) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bind_index), gdc.ConstTypePtr(pose.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Skin) GetBindPose(bind_index int, ) Transform3D {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bind_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965739696) // FIXME: should cache?
  var ret Transform3D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bind_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Skin) SetBindName(bind_index int, name StringName, )  {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bind_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3780747571) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bind_index), gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Skin) GetBindName(bind_index int, ) StringName {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bind_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 659327637) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bind_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Skin) SetBindBone(bind_index int, bone int, )  {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bind_bone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bind_index), gdc.ConstTypePtr(&bone), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Skin) GetBindBone(bind_index int, ) int {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bind_bone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bind_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Skin) ClearBinds()  {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_binds")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Signals
