// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForSkeletonModification2DFABRIKList struct {
  fnSetTargetNode gdc.MethodBindPtr
  fnGetTargetNode gdc.MethodBindPtr
  fnSetFabrikDataChainLength gdc.MethodBindPtr
  fnGetFabrikDataChainLength gdc.MethodBindPtr
  fnSetFabrikJointBone2DNode gdc.MethodBindPtr
  fnGetFabrikJointBone2DNode gdc.MethodBindPtr
  fnSetFabrikJointBoneIndex gdc.MethodBindPtr
  fnGetFabrikJointBoneIndex gdc.MethodBindPtr
  fnSetFabrikJointMagnetPosition gdc.MethodBindPtr
  fnGetFabrikJointMagnetPosition gdc.MethodBindPtr
  fnSetFabrikJointUseTargetRotation gdc.MethodBindPtr
  fnGetFabrikJointUseTargetRotation gdc.MethodBindPtr
}

var ptrsForSkeletonModification2DFABRIK ptrsForSkeletonModification2DFABRIKList

func initSkeletonModification2DFABRIKPtrs(iface gdc.Interface) {

  className := StringNameFromStr("SkeletonModification2DFABRIK")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_target_node")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DFABRIK.fnSetTargetNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
  }
  {
    methodName := StringNameFromStr("get_target_node")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DFABRIK.fnGetTargetNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
  }
  {
    methodName := StringNameFromStr("set_fabrik_data_chain_length")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DFABRIK.fnSetFabrikDataChainLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_fabrik_data_chain_length")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DFABRIK.fnGetFabrikDataChainLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_fabrik_joint_bone2d_node")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DFABRIK.fnSetFabrikJointBone2DNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2761262315))
  }
  {
    methodName := StringNameFromStr("get_fabrik_joint_bone2d_node")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DFABRIK.fnGetFabrikJointBone2DNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 408788394))
  }
  {
    methodName := StringNameFromStr("set_fabrik_joint_bone_index")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DFABRIK.fnSetFabrikJointBoneIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
  }
  {
    methodName := StringNameFromStr("get_fabrik_joint_bone_index")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DFABRIK.fnGetFabrikJointBoneIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
  }
  {
    methodName := StringNameFromStr("set_fabrik_joint_magnet_position")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DFABRIK.fnSetFabrikJointMagnetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 163021252))
  }
  {
    methodName := StringNameFromStr("get_fabrik_joint_magnet_position")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DFABRIK.fnGetFabrikJointMagnetPosition = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2299179447))
  }
  {
    methodName := StringNameFromStr("set_fabrik_joint_use_target_rotation")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DFABRIK.fnSetFabrikJointUseTargetRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("get_fabrik_joint_use_target_rotation")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DFABRIK.fnGetFabrikJointUseTargetRotation = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
}

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
  cargs := []gdc.ConstTypePtr{target_nodepath.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DFABRIK.fnSetTargetNode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DFABRIK) GetTargetNode() NodePath {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DFABRIK.fnGetTargetNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SkeletonModification2DFABRIK) SetFabrikDataChainLength(length int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DFABRIK.fnSetFabrikDataChainLength), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DFABRIK) GetFabrikDataChainLength() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DFABRIK.fnGetFabrikDataChainLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DFABRIK) SetFabrikJointBone2DNode(joint_idx int64, bone2d_nodepath NodePath, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , bone2d_nodepath.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DFABRIK.fnSetFabrikJointBone2DNode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DFABRIK) GetFabrikJointBone2DNode(joint_idx int64, ) NodePath {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()
  pinner.Pin(&joint_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DFABRIK.fnGetFabrikJointBone2DNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SkeletonModification2DFABRIK) SetFabrikJointBoneIndex(joint_idx int64, bone_idx int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DFABRIK.fnSetFabrikJointBoneIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DFABRIK) GetFabrikJointBoneIndex(joint_idx int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&joint_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DFABRIK.fnGetFabrikJointBoneIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DFABRIK) SetFabrikJointMagnetPosition(joint_idx int64, magnet_position Vector2, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , magnet_position.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DFABRIK.fnSetFabrikJointMagnetPosition), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DFABRIK) GetFabrikJointMagnetPosition(joint_idx int64, ) Vector2 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()
  pinner.Pin(&joint_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DFABRIK.fnGetFabrikJointMagnetPosition), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SkeletonModification2DFABRIK) SetFabrikJointUseTargetRotation(joint_idx int64, use_target_rotation bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , gdc.ConstTypePtr(&use_target_rotation) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DFABRIK.fnSetFabrikJointUseTargetRotation), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DFABRIK) GetFabrikJointUseTargetRotation(joint_idx int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&joint_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DFABRIK.fnGetFabrikJointUseTargetRotation), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
