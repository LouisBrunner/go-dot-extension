// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type BoneAttachment3D struct {
  obj gdc.ObjectPtr
}

func (me *BoneAttachment3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *BoneAttachment3D) BaseClass() string {
  return "BoneAttachment3D"
}



// Enums

func (me *BoneAttachment3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *BoneAttachment3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *BoneAttachment3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *BoneAttachment3D) SetBoneName(bone_name String, )  {
  panic("TODO: implement")
}

func  (me *BoneAttachment3D) GetBoneName()  {
  panic("TODO: implement")
}

func  (me *BoneAttachment3D) SetBoneIdx(bone_idx int, )  {
  panic("TODO: implement")
}

func  (me *BoneAttachment3D) GetBoneIdx()  {
  panic("TODO: implement")
}

func  (me *BoneAttachment3D) OnBonePoseUpdate(bone_index int, )  {
  panic("TODO: implement")
}

func  (me *BoneAttachment3D) SetOverridePose(override_pose bool, )  {
  panic("TODO: implement")
}

func  (me *BoneAttachment3D) GetOverridePose()  {
  panic("TODO: implement")
}

func  (me *BoneAttachment3D) SetUseExternalSkeleton(use_external_skeleton bool, )  {
  panic("TODO: implement")
}

func  (me *BoneAttachment3D) GetUseExternalSkeleton()  {
  panic("TODO: implement")
}

func  (me *BoneAttachment3D) SetExternalSkeleton(external_skeleton NodePath, )  {
  panic("TODO: implement")
}

func  (me *BoneAttachment3D) GetExternalSkeleton()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
