// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SkeletonModification2DTwoBoneIK struct {
  obj gdc.ObjectPtr
}

func (me *SkeletonModification2DTwoBoneIK) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SkeletonModification2DTwoBoneIK) BaseClass() string {
  return "SkeletonModification2DTwoBoneIK"
}



// Enums

func (me *SkeletonModification2DTwoBoneIK) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SkeletonModification2DTwoBoneIK) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SkeletonModification2DTwoBoneIK) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SkeletonModification2DTwoBoneIK) SetTargetNode(target_nodepath NodePath, )  {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_target_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(target_nodepath.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DTwoBoneIK) GetTargetNode() NodePath {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_target_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DTwoBoneIK) SetTargetMinimumDistance(minimum_distance float32, )  {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_target_minimum_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&minimum_distance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DTwoBoneIK) GetTargetMinimumDistance() float32 {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_target_minimum_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DTwoBoneIK) SetTargetMaximumDistance(maximum_distance float32, )  {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_target_maximum_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&maximum_distance), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DTwoBoneIK) GetTargetMaximumDistance() float32 {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_target_maximum_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DTwoBoneIK) SetFlipBendDirection(flip_direction bool, )  {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flip_bend_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_direction), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DTwoBoneIK) GetFlipBendDirection() bool {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_flip_bend_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DTwoBoneIK) SetJointOneBone2DNode(bone2d_node NodePath, )  {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_joint_one_bone2d_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bone2d_node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DTwoBoneIK) GetJointOneBone2DNode() NodePath {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joint_one_bone2d_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DTwoBoneIK) SetJointOneBoneIdx(bone_idx int, )  {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_joint_one_bone_idx")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DTwoBoneIK) GetJointOneBoneIdx() int {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joint_one_bone_idx")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DTwoBoneIK) SetJointTwoBone2DNode(bone2d_node NodePath, )  {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_joint_two_bone2d_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bone2d_node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DTwoBoneIK) GetJointTwoBone2DNode() NodePath {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joint_two_bone2d_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DTwoBoneIK) SetJointTwoBoneIdx(bone_idx int, )  {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_joint_two_bone_idx")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DTwoBoneIK) GetJointTwoBoneIdx() int {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joint_two_bone_idx")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *SkeletonModification2DTwoBoneIK) GetPropTargetNodepath() NodePath {
  panic("TODO: implement")
}

func (me *SkeletonModification2DTwoBoneIK) SetPropTargetNodepath(value NodePath) {
  panic("TODO: implement")
}

func (me *SkeletonModification2DTwoBoneIK) GetPropTargetMinimumDistance() float32 {
  panic("TODO: implement")
}

func (me *SkeletonModification2DTwoBoneIK) SetPropTargetMinimumDistance(value float32) {
  panic("TODO: implement")
}

func (me *SkeletonModification2DTwoBoneIK) GetPropTargetMaximumDistance() float32 {
  panic("TODO: implement")
}

func (me *SkeletonModification2DTwoBoneIK) SetPropTargetMaximumDistance(value float32) {
  panic("TODO: implement")
}

func (me *SkeletonModification2DTwoBoneIK) GetPropFlipBendDirection() bool {
  panic("TODO: implement")
}

func (me *SkeletonModification2DTwoBoneIK) SetPropFlipBendDirection(value bool) {
  panic("TODO: implement")
}