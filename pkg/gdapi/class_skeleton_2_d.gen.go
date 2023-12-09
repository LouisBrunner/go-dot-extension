// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Skeleton2D struct {
  obj gdc.ObjectPtr
}

func (me *Skeleton2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Skeleton2D) BaseClass() string {
  return "Skeleton2D"
}



// Enums

func (me *Skeleton2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Skeleton2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Skeleton2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Skeleton2D) GetBoneCount()  {
  panic("TODO: implement")
}

func  (me *Skeleton2D) GetBone(idx int, )  {
  panic("TODO: implement")
}

func  (me *Skeleton2D) GetSkeleton()  {
  panic("TODO: implement")
}

func  (me *Skeleton2D) SetModificationStack(modification_stack SkeletonModificationStack2D, )  {
  panic("TODO: implement")
}

func  (me *Skeleton2D) GetModificationStack()  {
  panic("TODO: implement")
}

func  (me *Skeleton2D) ExecuteModifications(delta float32, execution_mode int, )  {
  panic("TODO: implement")
}

func  (me *Skeleton2D) SetBoneLocalPoseOverride(bone_idx int, override_pose Transform2D, strength float32, persistent bool, )  {
  panic("TODO: implement")
}

func  (me *Skeleton2D) GetBoneLocalPoseOverride(bone_idx int, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
