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

type SkeletonModification2DJiggle struct {
  SkeletonModification2D
}

func (me *SkeletonModification2DJiggle) BaseClass() string {
  return "SkeletonModification2DJiggle"
}

func NewSkeletonModification2DJiggle() *SkeletonModification2DJiggle {
  str := StringNameFromStr("SkeletonModification2DJiggle") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SkeletonModification2DJiggle{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *SkeletonModification2DJiggle) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SkeletonModification2DJiggle) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SkeletonModification2DJiggle) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SkeletonModification2DJiggle) SetTargetNode(target_nodepath NodePath, )  {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_target_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{target_nodepath.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DJiggle) GetTargetNode() NodePath {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
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

func  (me *SkeletonModification2DJiggle) SetJiggleDataChainLength(length int64, )  {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_jiggle_data_chain_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DJiggle) GetJiggleDataChainLength() int64 {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_jiggle_data_chain_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DJiggle) SetStiffness(stiffness float64, )  {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stiffness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stiffness) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DJiggle) GetStiffness() float64 {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stiffness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DJiggle) SetMass(mass float64, )  {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mass) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DJiggle) GetMass() float64 {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DJiggle) SetDamping(damping float64, )  {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_damping")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&damping) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DJiggle) GetDamping() float64 {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_damping")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DJiggle) SetUseGravity(use_gravity bool, )  {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_gravity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DJiggle) GetUseGravity() bool {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DJiggle) SetGravity(gravity Vector2, )  {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gravity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DJiggle) GetGravity() Vector2 {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SkeletonModification2DJiggle) SetUseColliders(use_colliders bool, )  {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_use_colliders")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&use_colliders) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DJiggle) GetUseColliders() bool {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_use_colliders")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DJiggle) SetCollisionMask(collision_mask int64, )  {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&collision_mask) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DJiggle) GetCollisionMask() int64 {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_collision_mask")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DJiggle) SetJiggleJointBone2DNode(joint_idx int64, bone2d_node NodePath, )  {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_jiggle_joint_bone2d_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2761262315) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , bone2d_node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DJiggle) GetJiggleJointBone2DNode(joint_idx int64, ) NodePath {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_jiggle_joint_bone2d_node")
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

func  (me *SkeletonModification2DJiggle) SetJiggleJointBoneIndex(joint_idx int64, bone_idx int64, )  {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_jiggle_joint_bone_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3937882851) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DJiggle) GetJiggleJointBoneIndex(joint_idx int64, ) int64 {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_jiggle_joint_bone_index")
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

func  (me *SkeletonModification2DJiggle) SetJiggleJointOverride(joint_idx int64, override bool, )  {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_jiggle_joint_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , gdc.ConstTypePtr(&override) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DJiggle) GetJiggleJointOverride(joint_idx int64, ) bool {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_jiggle_joint_override")
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

func  (me *SkeletonModification2DJiggle) SetJiggleJointStiffness(joint_idx int64, stiffness float64, )  {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_jiggle_joint_stiffness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , gdc.ConstTypePtr(&stiffness) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DJiggle) GetJiggleJointStiffness(joint_idx int64, ) float64 {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_jiggle_joint_stiffness")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&joint_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DJiggle) SetJiggleJointMass(joint_idx int64, mass float64, )  {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_jiggle_joint_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , gdc.ConstTypePtr(&mass) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DJiggle) GetJiggleJointMass(joint_idx int64, ) float64 {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_jiggle_joint_mass")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&joint_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DJiggle) SetJiggleJointDamping(joint_idx int64, damping float64, )  {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_jiggle_joint_damping")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , gdc.ConstTypePtr(&damping) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DJiggle) GetJiggleJointDamping(joint_idx int64, ) float64 {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_jiggle_joint_damping")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&joint_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DJiggle) SetJiggleJointUseGravity(joint_idx int64, use_gravity bool, )  {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_jiggle_joint_use_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 300928843) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , gdc.ConstTypePtr(&use_gravity) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DJiggle) GetJiggleJointUseGravity(joint_idx int64, ) bool {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_jiggle_joint_use_gravity")
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

func  (me *SkeletonModification2DJiggle) SetJiggleJointGravity(joint_idx int64, gravity Vector2, )  {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_jiggle_joint_gravity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 163021252) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , gravity.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DJiggle) GetJiggleJointGravity(joint_idx int64, ) Vector2 {
  classNameV := StringNameFromStr("SkeletonModification2DJiggle")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_jiggle_joint_gravity")
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
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
