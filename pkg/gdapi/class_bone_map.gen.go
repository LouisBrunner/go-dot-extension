// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BoneMap struct {
  obj gdc.ObjectPtr
}

func (me *BoneMap) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *BoneMap) BaseClass() string {
  return "BoneMap"
}



// Enums

func (me *BoneMap) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *BoneMap) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *BoneMap) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *BoneMap) GetProfile()  {
  panic("TODO: implement")
}

func  (me *BoneMap) SetProfile(profile SkeletonProfile, )  {
  panic("TODO: implement")
}

func  (me *BoneMap) GetSkeletonBoneName(profile_bone_name StringName, )  {
  panic("TODO: implement")
}

func  (me *BoneMap) SetSkeletonBoneName(profile_bone_name StringName, skeleton_bone_name StringName, )  {
  panic("TODO: implement")
}

func  (me *BoneMap) FindProfileBoneName(skeleton_bone_name StringName, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
