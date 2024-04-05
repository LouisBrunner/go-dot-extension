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

type ptrsForPhysicalBone2DList struct {
  fnGetJoint gdc.MethodBindPtr
  fnGetAutoConfigureJoint gdc.MethodBindPtr
  fnSetAutoConfigureJoint gdc.MethodBindPtr
  fnSetSimulatePhysics gdc.MethodBindPtr
  fnGetSimulatePhysics gdc.MethodBindPtr
  fnIsSimulatingPhysics gdc.MethodBindPtr
  fnSetBone2DNodepath gdc.MethodBindPtr
  fnGetBone2DNodepath gdc.MethodBindPtr
  fnSetBone2DIndex gdc.MethodBindPtr
  fnGetBone2DIndex gdc.MethodBindPtr
  fnSetFollowBoneWhenSimulating gdc.MethodBindPtr
  fnGetFollowBoneWhenSimulating gdc.MethodBindPtr
}

var ptrsForPhysicalBone2D ptrsForPhysicalBone2DList

func initPhysicalBone2DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("PhysicalBone2D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_joint")
    defer methodName.Destroy()
    ptrsForPhysicalBone2D.fnGetJoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3582132112))
  }
  {
    methodName := StringNameFromStr("get_auto_configure_joint")
    defer methodName.Destroy()
    ptrsForPhysicalBone2D.fnGetAutoConfigureJoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_auto_configure_joint")
    defer methodName.Destroy()
    ptrsForPhysicalBone2D.fnSetAutoConfigureJoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("set_simulate_physics")
    defer methodName.Destroy()
    ptrsForPhysicalBone2D.fnSetSimulatePhysics = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_simulate_physics")
    defer methodName.Destroy()
    ptrsForPhysicalBone2D.fnGetSimulatePhysics = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("is_simulating_physics")
    defer methodName.Destroy()
    ptrsForPhysicalBone2D.fnIsSimulatingPhysics = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_bone2d_nodepath")
    defer methodName.Destroy()
    ptrsForPhysicalBone2D.fnSetBone2DNodepath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
  }
  {
    methodName := StringNameFromStr("get_bone2d_nodepath")
    defer methodName.Destroy()
    ptrsForPhysicalBone2D.fnGetBone2DNodepath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
  }
  {
    methodName := StringNameFromStr("set_bone2d_index")
    defer methodName.Destroy()
    ptrsForPhysicalBone2D.fnSetBone2DIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_bone2d_index")
    defer methodName.Destroy()
    ptrsForPhysicalBone2D.fnGetBone2DIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_follow_bone_when_simulating")
    defer methodName.Destroy()
    ptrsForPhysicalBone2D.fnSetFollowBoneWhenSimulating = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_follow_bone_when_simulating")
    defer methodName.Destroy()
    ptrsForPhysicalBone2D.fnGetFollowBoneWhenSimulating = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

type PhysicalBone2D struct {
  RigidBody2D
}

func (me *PhysicalBone2D) BaseClass() string {
  return "PhysicalBone2D"
}

func NewPhysicalBone2D() *PhysicalBone2D {
  str := StringNameFromStr("PhysicalBone2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PhysicalBone2D{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewJoint2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone2D.fnGetJoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicalBone2D) GetAutoConfigureJoint() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone2D.fnGetAutoConfigureJoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicalBone2D) SetAutoConfigureJoint(auto_configure_joint bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&auto_configure_joint) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone2D.fnSetAutoConfigureJoint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicalBone2D) SetSimulatePhysics(simulate_physics bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&simulate_physics) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone2D.fnSetSimulatePhysics), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicalBone2D) GetSimulatePhysics() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone2D.fnGetSimulatePhysics), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicalBone2D) IsSimulatingPhysics() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone2D.fnIsSimulatingPhysics), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicalBone2D) SetBone2DNodepath(nodepath NodePath, )  {
  cargs := []gdc.ConstTypePtr{nodepath.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone2D.fnSetBone2DNodepath), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicalBone2D) GetBone2DNodepath() NodePath {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone2D.fnGetBone2DNodepath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicalBone2D) SetBone2DIndex(bone_index int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone2D.fnSetBone2DIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicalBone2D) GetBone2DIndex() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone2D.fnGetBone2DIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PhysicalBone2D) SetFollowBoneWhenSimulating(follow_bone bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&follow_bone) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone2D.fnSetFollowBoneWhenSimulating), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PhysicalBone2D) GetFollowBoneWhenSimulating() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicalBone2D.fnGetFollowBoneWhenSimulating), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
