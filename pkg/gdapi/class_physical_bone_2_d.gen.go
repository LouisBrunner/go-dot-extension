// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func  (me *PhysicalBone2D) GetJoint()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone2D) GetAutoConfigureJoint()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone2D) SetAutoConfigureJoint(auto_configure_joint bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone2D) SetSimulatePhysics(simulate_physics bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone2D) GetSimulatePhysics()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone2D) IsSimulatingPhysics()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone2D) SetBone2DNodepath(nodepath NodePath, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone2D) GetBone2DNodepath()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone2D) SetBone2DIndex(bone_index int, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone2D) GetBone2DIndex()  {
  panic("TODO: implement")
}

func  (me *PhysicalBone2D) SetFollowBoneWhenSimulating(follow_bone bool, )  {
  panic("TODO: implement")
}

func  (me *PhysicalBone2D) GetFollowBoneWhenSimulating()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
