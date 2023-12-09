// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type SkeletonIK3D struct {
  obj gdc.ObjectPtr
}

func (me *SkeletonIK3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SkeletonIK3D) BaseClass() string {
  return "SkeletonIK3D"
}



// Enums

func (me *SkeletonIK3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SkeletonIK3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SkeletonIK3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *SkeletonIK3D) SetRootBone(root_bone StringName, )  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) GetRootBone()  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) SetTipBone(tip_bone StringName, )  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) GetTipBone()  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) SetInterpolation(interpolation float32, )  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) GetInterpolation()  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) SetTargetTransform(target Transform3D, )  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) GetTargetTransform()  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) SetTargetNode(node NodePath, )  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) GetTargetNode()  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) SetOverrideTipBasis(override bool, )  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) IsOverrideTipBasis()  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) SetUseMagnet(use bool, )  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) IsUsingMagnet()  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) SetMagnetPosition(local_position Vector3, )  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) GetMagnetPosition()  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) GetParentSkeleton()  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) IsRunning()  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) SetMinDistance(min_distance float32, )  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) GetMinDistance()  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) SetMaxIterations(iterations int, )  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) GetMaxIterations()  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) Start(one_time bool, )  {
  panic("TODO: implement")
}

func  (me *SkeletonIK3D) Stop()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
