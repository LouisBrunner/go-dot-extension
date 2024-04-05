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

type ptrsForSkeletonModification2DTwoBoneIKList struct {
  fnSetTargetNode gdc.MethodBindPtr
  fnGetTargetNode gdc.MethodBindPtr
  fnSetTargetMinimumDistance gdc.MethodBindPtr
  fnGetTargetMinimumDistance gdc.MethodBindPtr
  fnSetTargetMaximumDistance gdc.MethodBindPtr
  fnGetTargetMaximumDistance gdc.MethodBindPtr
  fnSetFlipBendDirection gdc.MethodBindPtr
  fnGetFlipBendDirection gdc.MethodBindPtr
  fnSetJointOneBone2DNode gdc.MethodBindPtr
  fnGetJointOneBone2DNode gdc.MethodBindPtr
  fnSetJointOneBoneIdx gdc.MethodBindPtr
  fnGetJointOneBoneIdx gdc.MethodBindPtr
  fnSetJointTwoBone2DNode gdc.MethodBindPtr
  fnGetJointTwoBone2DNode gdc.MethodBindPtr
  fnSetJointTwoBoneIdx gdc.MethodBindPtr
  fnGetJointTwoBoneIdx gdc.MethodBindPtr
}

var ptrsForSkeletonModification2DTwoBoneIK ptrsForSkeletonModification2DTwoBoneIKList

func initSkeletonModification2DTwoBoneIKPtrs(iface gdc.Interface) {

  className := StringNameFromStr("SkeletonModification2DTwoBoneIK")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_target_node")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DTwoBoneIK.fnSetTargetNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
  }
  {
    methodName := StringNameFromStr("get_target_node")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DTwoBoneIK.fnGetTargetNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
  }
  {
    methodName := StringNameFromStr("set_target_minimum_distance")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DTwoBoneIK.fnSetTargetMinimumDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_target_minimum_distance")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DTwoBoneIK.fnGetTargetMinimumDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_target_maximum_distance")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DTwoBoneIK.fnSetTargetMaximumDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 373806689))
  }
  {
    methodName := StringNameFromStr("get_target_maximum_distance")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DTwoBoneIK.fnGetTargetMaximumDistance = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1740695150))
  }
  {
    methodName := StringNameFromStr("set_flip_bend_direction")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DTwoBoneIK.fnSetFlipBendDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_flip_bend_direction")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DTwoBoneIK.fnGetFlipBendDirection = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_joint_one_bone2d_node")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DTwoBoneIK.fnSetJointOneBone2DNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
  }
  {
    methodName := StringNameFromStr("get_joint_one_bone2d_node")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DTwoBoneIK.fnGetJointOneBone2DNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
  }
  {
    methodName := StringNameFromStr("set_joint_one_bone_idx")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DTwoBoneIK.fnSetJointOneBoneIdx = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_joint_one_bone_idx")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DTwoBoneIK.fnGetJointOneBoneIdx = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("set_joint_two_bone2d_node")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DTwoBoneIK.fnSetJointTwoBone2DNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1348162250))
  }
  {
    methodName := StringNameFromStr("get_joint_two_bone2d_node")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DTwoBoneIK.fnGetJointTwoBone2DNode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 4075236667))
  }
  {
    methodName := StringNameFromStr("set_joint_two_bone_idx")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DTwoBoneIK.fnSetJointTwoBoneIdx = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_joint_two_bone_idx")
    defer methodName.Destroy()
    ptrsForSkeletonModification2DTwoBoneIK.fnGetJointTwoBoneIdx = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
}

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
  cargs := []gdc.ConstTypePtr{target_nodepath.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DTwoBoneIK.fnSetTargetNode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DTwoBoneIK) GetTargetNode() NodePath {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DTwoBoneIK.fnGetTargetNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SkeletonModification2DTwoBoneIK) SetTargetMinimumDistance(minimum_distance float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&minimum_distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DTwoBoneIK.fnSetTargetMinimumDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DTwoBoneIK) GetTargetMinimumDistance() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DTwoBoneIK.fnGetTargetMinimumDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DTwoBoneIK) SetTargetMaximumDistance(maximum_distance float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&maximum_distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DTwoBoneIK.fnSetTargetMaximumDistance), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DTwoBoneIK) GetTargetMaximumDistance() float64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DTwoBoneIK.fnGetTargetMaximumDistance), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DTwoBoneIK) SetFlipBendDirection(flip_direction bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&flip_direction) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DTwoBoneIK.fnSetFlipBendDirection), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DTwoBoneIK) GetFlipBendDirection() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DTwoBoneIK.fnGetFlipBendDirection), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DTwoBoneIK) SetJointOneBone2DNode(bone2d_node NodePath, )  {
  cargs := []gdc.ConstTypePtr{bone2d_node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DTwoBoneIK.fnSetJointOneBone2DNode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DTwoBoneIK) GetJointOneBone2DNode() NodePath {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DTwoBoneIK.fnGetJointOneBone2DNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SkeletonModification2DTwoBoneIK) SetJointOneBoneIdx(bone_idx int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DTwoBoneIK.fnSetJointOneBoneIdx), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DTwoBoneIK) GetJointOneBoneIdx() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DTwoBoneIK.fnGetJointOneBoneIdx), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DTwoBoneIK) SetJointTwoBone2DNode(bone2d_node NodePath, )  {
  cargs := []gdc.ConstTypePtr{bone2d_node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DTwoBoneIK.fnSetJointTwoBone2DNode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DTwoBoneIK) GetJointTwoBone2DNode() NodePath {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewNodePath()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DTwoBoneIK.fnGetJointTwoBone2DNode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SkeletonModification2DTwoBoneIK) SetJointTwoBoneIdx(bone_idx int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DTwoBoneIK.fnSetJointTwoBoneIdx), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DTwoBoneIK) GetJointTwoBoneIdx() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2DTwoBoneIK.fnGetJointTwoBoneIdx), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
