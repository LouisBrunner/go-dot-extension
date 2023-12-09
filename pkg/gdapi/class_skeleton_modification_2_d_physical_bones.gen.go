// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SkeletonModification2DPhysicalBones struct {
  obj gdc.ObjectPtr
}

func (me *SkeletonModification2DPhysicalBones) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SkeletonModification2DPhysicalBones) BaseClass() string {
  return "SkeletonModification2DPhysicalBones"
}



// Enums

func (me *SkeletonModification2DPhysicalBones) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SkeletonModification2DPhysicalBones) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *SkeletonModification2DPhysicalBones) SetPhysicalBoneChainLength(length int, )  {
  panic("TODO: implement")
}

func  (me *SkeletonModification2DPhysicalBones) GetPhysicalBoneChainLength()  {
  panic("TODO: implement")
}

func  (me *SkeletonModification2DPhysicalBones) SetPhysicalBoneNode(joint_idx int, physicalbone2d_node NodePath, )  {
  panic("TODO: implement")
}

func  (me *SkeletonModification2DPhysicalBones) GetPhysicalBoneNode(joint_idx int, )  {
  panic("TODO: implement")
}

func  (me *SkeletonModification2DPhysicalBones) FetchPhysicalBones()  {
  panic("TODO: implement")
}

func  (me *SkeletonModification2DPhysicalBones) StartSimulation(bones StringName, )  {
  panic("TODO: implement")
}

func  (me *SkeletonModification2DPhysicalBones) StopSimulation(bones StringName, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
