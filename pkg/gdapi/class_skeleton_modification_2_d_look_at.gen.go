// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SkeletonModification2DLookAt struct {
  obj gdc.ObjectPtr
}

func (me *SkeletonModification2DLookAt) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SkeletonModification2DLookAt) BaseClass() string {
  return "SkeletonModification2DLookAt"
}



// Enums

func (me *SkeletonModification2DLookAt) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SkeletonModification2DLookAt) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SkeletonModification2DLookAt) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SkeletonModification2DLookAt) SetBone2DNode(bone2d_nodepath NodePath, )  {
  classNameV := StringNameFromStr("SkeletonModification2DLookAt")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone2d_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bone2d_nodepath.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DLookAt) GetBone2DNode() NodePath {
  classNameV := StringNameFromStr("SkeletonModification2DLookAt")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone2d_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DLookAt) SetBoneIndex(bone_idx int, )  {
  classNameV := StringNameFromStr("SkeletonModification2DLookAt")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DLookAt) GetBoneIndex() int {
  classNameV := StringNameFromStr("SkeletonModification2DLookAt")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_index")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DLookAt) SetTargetNode(target_nodepath NodePath, )  {
  classNameV := StringNameFromStr("SkeletonModification2DLookAt")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_target_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1348162250) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(target_nodepath.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DLookAt) GetTargetNode() NodePath {
  classNameV := StringNameFromStr("SkeletonModification2DLookAt")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_target_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4075236667) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DLookAt) SetAdditionalRotation(rotation float32, )  {
  classNameV := StringNameFromStr("SkeletonModification2DLookAt")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_additional_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&rotation), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DLookAt) GetAdditionalRotation() float32 {
  classNameV := StringNameFromStr("SkeletonModification2DLookAt")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_additional_rotation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DLookAt) SetEnableConstraint(enable_constraint bool, )  {
  classNameV := StringNameFromStr("SkeletonModification2DLookAt")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enable_constraint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable_constraint), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DLookAt) GetEnableConstraint() bool {
  classNameV := StringNameFromStr("SkeletonModification2DLookAt")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enable_constraint")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DLookAt) SetConstraintAngleMin(angle_min float32, )  {
  classNameV := StringNameFromStr("SkeletonModification2DLookAt")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constraint_angle_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angle_min), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DLookAt) GetConstraintAngleMin() float32 {
  classNameV := StringNameFromStr("SkeletonModification2DLookAt")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constraint_angle_min")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DLookAt) SetConstraintAngleMax(angle_max float32, )  {
  classNameV := StringNameFromStr("SkeletonModification2DLookAt")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constraint_angle_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angle_max), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DLookAt) GetConstraintAngleMax() float32 {
  classNameV := StringNameFromStr("SkeletonModification2DLookAt")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constraint_angle_max")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DLookAt) SetConstraintAngleInvert(invert bool, )  {
  classNameV := StringNameFromStr("SkeletonModification2DLookAt")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constraint_angle_invert")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&invert), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DLookAt) GetConstraintAngleInvert() bool {
  classNameV := StringNameFromStr("SkeletonModification2DLookAt")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constraint_angle_invert")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *SkeletonModification2DLookAt) GetPropBoneIndex() int {
  panic("TODO: implement")
}

func (me *SkeletonModification2DLookAt) SetPropBoneIndex(value int) {
  panic("TODO: implement")
}

func (me *SkeletonModification2DLookAt) GetPropBone2DNode() NodePath {
  panic("TODO: implement")
}

func (me *SkeletonModification2DLookAt) SetPropBone2DNode(value NodePath) {
  panic("TODO: implement")
}

func (me *SkeletonModification2DLookAt) GetPropTargetNodepath() NodePath {
  panic("TODO: implement")
}

func (me *SkeletonModification2DLookAt) SetPropTargetNodepath(value NodePath) {
  panic("TODO: implement")
}