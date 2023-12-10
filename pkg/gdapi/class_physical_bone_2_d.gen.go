// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PhysicalBone2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicalBone2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicalBone2D) BaseClass() string {
  return "PhysicalBone2D"
}



// Enums

func (me *PhysicalBone2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicalBone2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicalBone2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PhysicalBone2D) GetJoint() Joint2D {
  classNameV := StringNameFromStr("PhysicalBone2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3582132112) // FIXME: should cache?
  var ret Joint2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone2D) GetAutoConfigureJoint() bool {
  classNameV := StringNameFromStr("PhysicalBone2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_auto_configure_joint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone2D) SetAutoConfigureJoint(auto_configure_joint bool, )  {
  classNameV := StringNameFromStr("PhysicalBone2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_auto_configure_joint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&auto_configure_joint), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone2D) SetSimulatePhysics(simulate_physics bool, )  {
  classNameV := StringNameFromStr("PhysicalBone2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_simulate_physics")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&simulate_physics), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone2D) GetSimulatePhysics() bool {
  classNameV := StringNameFromStr("PhysicalBone2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_simulate_physics")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone2D) IsSimulatingPhysics() bool {
  classNameV := StringNameFromStr("PhysicalBone2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_simulating_physics")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone2D) SetBone2DNodepath(nodepath NodePath, )  {
  classNameV := StringNameFromStr("PhysicalBone2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone2d_nodepath")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(nodepath.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone2D) GetBone2DNodepath() NodePath {
  classNameV := StringNameFromStr("PhysicalBone2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone2d_nodepath")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone2D) SetBone2DIndex(bone_index int, )  {
  classNameV := StringNameFromStr("PhysicalBone2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone2d_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_index), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone2D) GetBone2DIndex() int {
  classNameV := StringNameFromStr("PhysicalBone2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone2d_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicalBone2D) SetFollowBoneWhenSimulating(follow_bone bool, )  {
  classNameV := StringNameFromStr("PhysicalBone2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_follow_bone_when_simulating")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&follow_bone), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *PhysicalBone2D) GetFollowBoneWhenSimulating() bool {
  classNameV := StringNameFromStr("PhysicalBone2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_follow_bone_when_simulating")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *PhysicalBone2D) GetPropBone2DNodepath() NodePath {
  panic("TODO: implement")
}

func (me *PhysicalBone2D) SetPropBone2DNodepath(value NodePath) {
  panic("TODO: implement")
}

func (me *PhysicalBone2D) GetPropBone2DIndex() int {
  panic("TODO: implement")
}

func (me *PhysicalBone2D) SetPropBone2DIndex(value int) {
  panic("TODO: implement")
}

func (me *PhysicalBone2D) GetPropAutoConfigureJoint() bool {
  panic("TODO: implement")
}

func (me *PhysicalBone2D) SetPropAutoConfigureJoint(value bool) {
  panic("TODO: implement")
}

func (me *PhysicalBone2D) GetPropSimulatePhysics() bool {
  panic("TODO: implement")
}

func (me *PhysicalBone2D) SetPropSimulatePhysics(value bool) {
  panic("TODO: implement")
}

func (me *PhysicalBone2D) GetPropFollowBoneWhenSimulating() bool {
  panic("TODO: implement")
}

func (me *PhysicalBone2D) SetPropFollowBoneWhenSimulating(value bool) {
  panic("TODO: implement")
}