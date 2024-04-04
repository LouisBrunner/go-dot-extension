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

type SkeletonModification2DFABRIK struct {
  SkeletonModification2D
}

func (me *SkeletonModification2DFABRIK) BaseClass() string {
  return "SkeletonModification2DFABRIK"
}

func NewSkeletonModification2DFABRIK() *SkeletonModification2DFABRIK {
  str := StringNameFromStr("SkeletonModification2DFABRIK") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SkeletonModification2DFABRIK{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *SkeletonModification2DFABRIK) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SkeletonModification2DFABRIK) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SkeletonModification2DFABRIK) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SkeletonModification2DFABRIK) SetTargetNode(target_nodepath NodePath, )  {
  classNameV := StringNameFromStr("SkeletonModification2DFABRIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_target_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{target_nodepath.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DFABRIK) GetTargetNode() NodePath {
  classNameV := StringNameFromStr("SkeletonModification2DFABRIK")
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

func  (me *SkeletonModification2DFABRIK) SetFabrikDataChainLength(length int64, )  {
  classNameV := StringNameFromStr("SkeletonModification2DFABRIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fabrik_data_chain_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DFABRIK) GetFabrikDataChainLength() int64 {
  classNameV := StringNameFromStr("SkeletonModification2DFABRIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fabrik_data_chain_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DFABRIK) SetFabrikJointBone2DNode(joint_idx int64, bone2d_nodepath NodePath, )  {
  classNameV := StringNameFromStr("SkeletonModification2DFABRIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fabrik_joint_bone2d_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2761262315) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , bone2d_nodepath.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DFABRIK) GetFabrikJointBone2DNode(joint_idx int64, ) NodePath {
  classNameV := StringNameFromStr("SkeletonModification2DFABRIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fabrik_joint_bone2d_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 408788394) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()
  pinner.Pin(&joint_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SkeletonModification2DFABRIK) SetFabrikJointBoneIndex(joint_idx int64, bone_idx int64, )  {
  classNameV := StringNameFromStr("SkeletonModification2DFABRIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fabrik_joint_bone_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DFABRIK) GetFabrikJointBoneIndex(joint_idx int64, ) int64 {
  classNameV := StringNameFromStr("SkeletonModification2DFABRIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fabrik_joint_bone_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 923996154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&joint_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DFABRIK) SetFabrikJointMagnetPosition(joint_idx int64, magnet_position Vector2, )  {
  classNameV := StringNameFromStr("SkeletonModification2DFABRIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fabrik_joint_magnet_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 163021252) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , magnet_position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DFABRIK) GetFabrikJointMagnetPosition(joint_idx int64, ) Vector2 {
  classNameV := StringNameFromStr("SkeletonModification2DFABRIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fabrik_joint_magnet_position")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2299179447) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&joint_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SkeletonModification2DFABRIK) SetFabrikJointUseTargetRotation(joint_idx int64, use_target_rotation bool, )  {
  classNameV := StringNameFromStr("SkeletonModification2DFABRIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_fabrik_joint_use_target_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , gdc.ConstTypePtr(&use_target_rotation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DFABRIK) GetFabrikJointUseTargetRotation(joint_idx int64, ) bool {
  classNameV := StringNameFromStr("SkeletonModification2DFABRIK")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_fabrik_joint_use_target_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1116898809) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&joint_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
