// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type BoneAttachment3D struct {
  Node3D
}

func (me *BoneAttachment3D) BaseClass() string {
  return "BoneAttachment3D"
}

func NewBoneAttachment3D() *BoneAttachment3D {
  str := StringNameFromStr("BoneAttachment3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &BoneAttachment3D{}
  obj.SetBaseObject(objPtr)
  return obj
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
  classNameV := StringNameFromStr("BoneAttachment3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bone_name.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BoneAttachment3D) GetBoneName() String {
  classNameV := StringNameFromStr("BoneAttachment3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *BoneAttachment3D) SetBoneIdx(bone_idx int64, )  {
  classNameV := StringNameFromStr("BoneAttachment3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone_idx")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BoneAttachment3D) GetBoneIdx() int64 {
  classNameV := StringNameFromStr("BoneAttachment3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_idx")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BoneAttachment3D) OnBonePoseUpdate(bone_index int64, )  {
  classNameV := StringNameFromStr("BoneAttachment3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("on_bone_pose_update")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_index), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BoneAttachment3D) SetOverridePose(override_pose bool, )  {
  classNameV := StringNameFromStr("BoneAttachment3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_override_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&override_pose), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BoneAttachment3D) GetOverridePose() bool {
  classNameV := StringNameFromStr("BoneAttachment3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_override_pose")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BoneAttachment3D) SetUseExternalSkeleton(use_external_skeleton bool, )  {
  classNameV := StringNameFromStr("BoneAttachment3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_external_skeleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_external_skeleton), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BoneAttachment3D) GetUseExternalSkeleton() bool {
  classNameV := StringNameFromStr("BoneAttachment3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_external_skeleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *BoneAttachment3D) SetExternalSkeleton(external_skeleton NodePath, )  {
  classNameV := StringNameFromStr("BoneAttachment3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_external_skeleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(external_skeleton.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *BoneAttachment3D) GetExternalSkeleton() NodePath {
  classNameV := StringNameFromStr("BoneAttachment3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_external_skeleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
