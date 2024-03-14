// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SkeletonModification2DCCDIK struct {
  SkeletonModification2D
}

func (me *SkeletonModification2DCCDIK) BaseClass() string {
  return "SkeletonModification2DCCDIK"
}



// Enums

func (me *SkeletonModification2DCCDIK) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SkeletonModification2DCCDIK) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SkeletonModification2DCCDIK) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SkeletonModification2DCCDIK) SetTargetNode(target_nodepath NodePath, )  {
  classNameV := StringNameFromStr("SkeletonModification2DCCDIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_target_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(target_nodepath.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DCCDIK) GetTargetNode() NodePath {
  classNameV := StringNameFromStr("SkeletonModification2DCCDIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_target_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DCCDIK) SetTipNode(tip_nodepath NodePath, )  {
  classNameV := StringNameFromStr("SkeletonModification2DCCDIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_tip_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(tip_nodepath.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DCCDIK) GetTipNode() NodePath {
  classNameV := StringNameFromStr("SkeletonModification2DCCDIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_tip_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DCCDIK) SetCcdikDataChainLength(length int, )  {
  classNameV := StringNameFromStr("SkeletonModification2DCCDIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ccdik_data_chain_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DCCDIK) GetCcdikDataChainLength() int {
  classNameV := StringNameFromStr("SkeletonModification2DCCDIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ccdik_data_chain_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DCCDIK) SetCcdikJointBone2DNode(joint_idx int, bone2d_nodepath NodePath, )  {
  classNameV := StringNameFromStr("SkeletonModification2DCCDIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ccdik_joint_bone2d_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2761262315) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), gdc.ConstTypePtr(bone2d_nodepath.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DCCDIK) GetCcdikJointBone2DNode(joint_idx int, ) NodePath {
  classNameV := StringNameFromStr("SkeletonModification2DCCDIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ccdik_joint_bone2d_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 408788394) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DCCDIK) SetCcdikJointBoneIndex(joint_idx int, bone_idx int, )  {
  classNameV := StringNameFromStr("SkeletonModification2DCCDIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ccdik_joint_bone_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), gdc.ConstTypePtr(&bone_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DCCDIK) GetCcdikJointBoneIndex(joint_idx int, ) int {
  classNameV := StringNameFromStr("SkeletonModification2DCCDIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ccdik_joint_bone_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DCCDIK) SetCcdikJointRotateFromJoint(joint_idx int, rotate_from_joint bool, )  {
  classNameV := StringNameFromStr("SkeletonModification2DCCDIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ccdik_joint_rotate_from_joint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), gdc.ConstTypePtr(&rotate_from_joint), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DCCDIK) GetCcdikJointRotateFromJoint(joint_idx int, ) bool {
  classNameV := StringNameFromStr("SkeletonModification2DCCDIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ccdik_joint_rotate_from_joint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DCCDIK) SetCcdikJointEnableConstraint(joint_idx int, enable_constraint bool, )  {
  classNameV := StringNameFromStr("SkeletonModification2DCCDIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ccdik_joint_enable_constraint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), gdc.ConstTypePtr(&enable_constraint), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DCCDIK) GetCcdikJointEnableConstraint(joint_idx int, ) bool {
  classNameV := StringNameFromStr("SkeletonModification2DCCDIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ccdik_joint_enable_constraint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DCCDIK) SetCcdikJointConstraintAngleMin(joint_idx int, angle_min float32, )  {
  classNameV := StringNameFromStr("SkeletonModification2DCCDIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ccdik_joint_constraint_angle_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), gdc.ConstTypePtr(&angle_min), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DCCDIK) GetCcdikJointConstraintAngleMin(joint_idx int, ) float32 {
  classNameV := StringNameFromStr("SkeletonModification2DCCDIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ccdik_joint_constraint_angle_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DCCDIK) SetCcdikJointConstraintAngleMax(joint_idx int, angle_max float32, )  {
  classNameV := StringNameFromStr("SkeletonModification2DCCDIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ccdik_joint_constraint_angle_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), gdc.ConstTypePtr(&angle_max), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DCCDIK) GetCcdikJointConstraintAngleMax(joint_idx int, ) float32 {
  classNameV := StringNameFromStr("SkeletonModification2DCCDIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ccdik_joint_constraint_angle_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DCCDIK) SetCcdikJointConstraintAngleInvert(joint_idx int, invert bool, )  {
  classNameV := StringNameFromStr("SkeletonModification2DCCDIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_ccdik_joint_constraint_angle_invert")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), gdc.ConstTypePtr(&invert), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DCCDIK) GetCcdikJointConstraintAngleInvert(joint_idx int, ) bool {
  classNameV := StringNameFromStr("SkeletonModification2DCCDIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ccdik_joint_constraint_angle_invert")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
