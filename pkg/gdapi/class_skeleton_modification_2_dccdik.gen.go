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

type ptrsForSkeletonModification2DCCDIKList struct {
  fnSetTargetNode gdc.MethodBindPtr
  fnGetTargetNode gdc.MethodBindPtr
  fnSetTipNode gdc.MethodBindPtr
  fnGetTipNode gdc.MethodBindPtr
  fnSetCcdikDataChainLength gdc.MethodBindPtr
  fnGetCcdikDataChainLength gdc.MethodBindPtr
  fnSetCcdikJointBone2DNode gdc.MethodBindPtr
  fnGetCcdikJointBone2DNode gdc.MethodBindPtr
  fnSetCcdikJointBoneIndex gdc.MethodBindPtr
  fnGetCcdikJointBoneIndex gdc.MethodBindPtr
  fnSetCcdikJointRotateFromJoint gdc.MethodBindPtr
  fnGetCcdikJointRotateFromJoint gdc.MethodBindPtr
  fnSetCcdikJointEnableConstraint gdc.MethodBindPtr
  fnGetCcdikJointEnableConstraint gdc.MethodBindPtr
  fnSetCcdikJointConstraintAngleMin gdc.MethodBindPtr
  fnGetCcdikJointConstraintAngleMin gdc.MethodBindPtr
  fnSetCcdikJointConstraintAngleMax gdc.MethodBindPtr
  fnGetCcdikJointConstraintAngleMax gdc.MethodBindPtr
  fnSetCcdikJointConstraintAngleInvert gdc.MethodBindPtr
  fnGetCcdikJointConstraintAngleInvert gdc.MethodBindPtr
}

var ptrsForSkeletonModification2DCCDIK ptrsForSkeletonModification2DCCDIKList

func initSkeletonModification2DCCDIKPtrs(iface gdc.Interface) {

  className := StringNameFromStr("SkeletonModification2DCCDIK")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_target_node")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DCCDIK.fnSetTargetNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
  }
  {
    methodName := StringNameFromStr("get_target_node")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DCCDIK.fnGetTargetNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
  }
  {
    methodName := StringNameFromStr("set_tip_node")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DCCDIK.fnSetTipNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
  }
  {
    methodName := StringNameFromStr("get_tip_node")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DCCDIK.fnGetTipNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
  }
  {
    methodName := StringNameFromStr("set_ccdik_data_chain_length")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DCCDIK.fnSetCcdikDataChainLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_ccdik_data_chain_length")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DCCDIK.fnGetCcdikDataChainLength = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2455072627))
  }
  {
    methodName := StringNameFromStr("set_ccdik_joint_bone2d_node")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DCCDIK.fnSetCcdikJointBone2DNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2761262315))
  }
  {
    methodName := StringNameFromStr("get_ccdik_joint_bone2d_node")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DCCDIK.fnGetCcdikJointBone2DNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 408788394))
  }
  {
    methodName := StringNameFromStr("set_ccdik_joint_bone_index")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DCCDIK.fnSetCcdikJointBoneIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3937882851))
  }
  {
    methodName := StringNameFromStr("get_ccdik_joint_bone_index")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DCCDIK.fnGetCcdikJointBoneIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
  }
  {
    methodName := StringNameFromStr("set_ccdik_joint_rotate_from_joint")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DCCDIK.fnSetCcdikJointRotateFromJoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("get_ccdik_joint_rotate_from_joint")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DCCDIK.fnGetCcdikJointRotateFromJoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("set_ccdik_joint_enable_constraint")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DCCDIK.fnSetCcdikJointEnableConstraint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("get_ccdik_joint_enable_constraint")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DCCDIK.fnGetCcdikJointEnableConstraint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
  {
    methodName := StringNameFromStr("set_ccdik_joint_constraint_angle_min")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DCCDIK.fnSetCcdikJointConstraintAngleMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
  }
  {
    methodName := StringNameFromStr("get_ccdik_joint_constraint_angle_min")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DCCDIK.fnGetCcdikJointConstraintAngleMin = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
  }
  {
    methodName := StringNameFromStr("set_ccdik_joint_constraint_angle_max")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DCCDIK.fnSetCcdikJointConstraintAngleMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
  }
  {
    methodName := StringNameFromStr("get_ccdik_joint_constraint_angle_max")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DCCDIK.fnGetCcdikJointConstraintAngleMax = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
  }
  {
    methodName := StringNameFromStr("set_ccdik_joint_constraint_angle_invert")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DCCDIK.fnSetCcdikJointConstraintAngleInvert = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 300928843))
  }
  {
    methodName := StringNameFromStr("get_ccdik_joint_constraint_angle_invert")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DCCDIK.fnGetCcdikJointConstraintAngleInvert = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1116898809))
  }
}

type SkeletonModification2DCCDIK struct {
  SkeletonModification2D
}

func (me *SkeletonModification2DCCDIK) BaseClass() string {
  return "SkeletonModification2DCCDIK"
}

func NewSkeletonModification2DCCDIK() *SkeletonModification2DCCDIK {
  str := StringNameFromStr("SkeletonModification2DCCDIK") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SkeletonModification2DCCDIK{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{target_nodepath.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DCCDIK.fnSetTargetNode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DCCDIK) GetTargetNode() NodePath {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DCCDIK.fnGetTargetNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SkeletonModification2DCCDIK) SetTipNode(tip_nodepath NodePath, )  {
  cargs := []gdc.ConstTypePtr{tip_nodepath.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DCCDIK.fnSetTipNode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DCCDIK) GetTipNode() NodePath {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DCCDIK.fnGetTipNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SkeletonModification2DCCDIK) SetCcdikDataChainLength(length int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DCCDIK.fnSetCcdikDataChainLength), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DCCDIK) GetCcdikDataChainLength() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DCCDIK.fnGetCcdikDataChainLength), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DCCDIK) SetCcdikJointBone2DNode(joint_idx int64, bone2d_nodepath NodePath, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , bone2d_nodepath.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DCCDIK.fnSetCcdikJointBone2DNode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DCCDIK) GetCcdikJointBone2DNode(joint_idx int64, ) NodePath {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()
  pinner.Pin(&joint_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DCCDIK.fnGetCcdikJointBone2DNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SkeletonModification2DCCDIK) SetCcdikJointBoneIndex(joint_idx int64, bone_idx int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DCCDIK.fnSetCcdikJointBoneIndex), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DCCDIK) GetCcdikJointBoneIndex(joint_idx int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&joint_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DCCDIK.fnGetCcdikJointBoneIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DCCDIK) SetCcdikJointRotateFromJoint(joint_idx int64, rotate_from_joint bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , gdc.ConstTypePtr(&rotate_from_joint) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DCCDIK.fnSetCcdikJointRotateFromJoint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DCCDIK) GetCcdikJointRotateFromJoint(joint_idx int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&joint_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DCCDIK.fnGetCcdikJointRotateFromJoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DCCDIK) SetCcdikJointEnableConstraint(joint_idx int64, enable_constraint bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , gdc.ConstTypePtr(&enable_constraint) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DCCDIK.fnSetCcdikJointEnableConstraint), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DCCDIK) GetCcdikJointEnableConstraint(joint_idx int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&joint_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DCCDIK.fnGetCcdikJointEnableConstraint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DCCDIK) SetCcdikJointConstraintAngleMin(joint_idx int64, angle_min float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , gdc.ConstTypePtr(&angle_min) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DCCDIK.fnSetCcdikJointConstraintAngleMin), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DCCDIK) GetCcdikJointConstraintAngleMin(joint_idx int64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&joint_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DCCDIK.fnGetCcdikJointConstraintAngleMin), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DCCDIK) SetCcdikJointConstraintAngleMax(joint_idx int64, angle_max float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , gdc.ConstTypePtr(&angle_max) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DCCDIK.fnSetCcdikJointConstraintAngleMax), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DCCDIK) GetCcdikJointConstraintAngleMax(joint_idx int64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&joint_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DCCDIK.fnGetCcdikJointConstraintAngleMax), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DCCDIK) SetCcdikJointConstraintAngleInvert(joint_idx int64, invert bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , gdc.ConstTypePtr(&invert) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DCCDIK.fnSetCcdikJointConstraintAngleInvert), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DCCDIK) GetCcdikJointConstraintAngleInvert(joint_idx int64, ) bool {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()
  pinner.Pin(&joint_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DCCDIK.fnGetCcdikJointConstraintAngleInvert), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
