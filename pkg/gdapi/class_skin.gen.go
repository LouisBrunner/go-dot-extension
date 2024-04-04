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

type Skin struct {
  Resource
}

func (me *Skin) BaseClass() string {
  return "Skin"
}

func NewSkin() *Skin {
  str := StringNameFromStr("Skin") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Skin{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *Skin) SetBindCount(bind_count int64, )  {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bind_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bind_count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skin) GetBindCount() int64 {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bind_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Skin) AddBind(bone int64, pose Transform3D, )  {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_bind")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3616898986) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone) , pose.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skin) AddNamedBind(name String, pose Transform3D, )  {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_named_bind")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3154712474) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), pose.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skin) SetBindPose(bind_index int64, pose Transform3D, )  {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bind_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3616898986) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bind_index) , pose.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skin) GetBindPose(bind_index int64, ) Transform3D {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bind_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1965739696) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bind_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform3D()
  pinner.Pin(&bind_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skin) SetBindName(bind_index int64, name StringName, )  {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bind_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3780747571) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bind_index) , name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skin) GetBindName(bind_index int64, ) StringName {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bind_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 659327637) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bind_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()
  pinner.Pin(&bind_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skin) SetBindBone(bind_index int64, bone int64, )  {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bind_bone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bind_index) , gdc.ConstTypePtr(&bone) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skin) GetBindBone(bind_index int64, ) int64 {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bind_bone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bind_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&bind_index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Skin) ClearBinds()  {
  classNameV := StringNameFromStr("Skin")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clear_binds")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

// Signals
