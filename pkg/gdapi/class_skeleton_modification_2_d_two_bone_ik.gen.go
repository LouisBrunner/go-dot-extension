// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type SkeletonModification2DTwoBoneIK struct {
  SkeletonModification2D
}

func (me *SkeletonModification2DTwoBoneIK) BaseClass() string {
  return "SkeletonModification2DTwoBoneIK"
}

func NewSkeletonModification2DTwoBoneIK() *SkeletonModification2DTwoBoneIK {
  str := StringNameFromStr("SkeletonModification2DTwoBoneIK") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SkeletonModification2DTwoBoneIK{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{target_nodepath.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DTwoBoneIK) GetTargetNode() NodePath {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_target_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SkeletonModification2DTwoBoneIK) SetTargetMinimumDistance(minimum_distance float64, )  {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_target_minimum_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&minimum_distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DTwoBoneIK) GetTargetMinimumDistance() float64 {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_target_minimum_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DTwoBoneIK) SetTargetMaximumDistance(maximum_distance float64, )  {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_target_maximum_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&maximum_distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DTwoBoneIK) GetTargetMaximumDistance() float64 {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_target_maximum_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DTwoBoneIK) SetFlipBendDirection(flip_direction bool, )  {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_flip_bend_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_direction) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DTwoBoneIK) GetFlipBendDirection() bool {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_flip_bend_direction")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DTwoBoneIK) SetJointOneBone2DNode(bone2d_node NodePath, )  {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_joint_one_bone2d_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{bone2d_node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DTwoBoneIK) GetJointOneBone2DNode() NodePath {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joint_one_bone2d_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SkeletonModification2DTwoBoneIK) SetJointOneBoneIdx(bone_idx int64, )  {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_joint_one_bone_idx")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DTwoBoneIK) GetJointOneBoneIdx() int64 {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joint_one_bone_idx")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DTwoBoneIK) SetJointTwoBone2DNode(bone2d_node NodePath, )  {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_joint_two_bone2d_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{bone2d_node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DTwoBoneIK) GetJointTwoBone2DNode() NodePath {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joint_two_bone2d_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SkeletonModification2DTwoBoneIK) SetJointTwoBoneIdx(bone_idx int64, )  {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_joint_two_bone_idx")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DTwoBoneIK) GetJointTwoBoneIdx() int64 {
  classNameV := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_joint_two_bone_idx")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
