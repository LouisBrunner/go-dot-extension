// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SkeletonModification2DPhysicalBones struct {
  obj gdc.ObjectPtr
}

func (me *SkeletonModification2DPhysicalBones) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SkeletonModification2DPhysicalBones) BaseClass() string {
  return "SkeletonModification2DPhysicalBones"
}



// Enums

func (me *SkeletonModification2DPhysicalBones) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SkeletonModification2DPhysicalBones) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SkeletonModification2DPhysicalBones) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SkeletonModification2DPhysicalBones) SetPhysicalBoneChainLength(length int, )  {
  classNameV := StringNameFromStr("SkeletonModification2DPhysicalBones")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_physical_bone_chain_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DPhysicalBones) GetPhysicalBoneChainLength() int {
  classNameV := StringNameFromStr("SkeletonModification2DPhysicalBones")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_physical_bone_chain_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DPhysicalBones) SetPhysicalBoneNode(joint_idx int, physicalbone2d_node NodePath, )  {
  classNameV := StringNameFromStr("SkeletonModification2DPhysicalBones")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_physical_bone_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2761262315) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), gdc.ConstTypePtr(physicalbone2d_node.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DPhysicalBones) GetPhysicalBoneNode(joint_idx int, ) NodePath {
  classNameV := StringNameFromStr("SkeletonModification2DPhysicalBones")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_physical_bone_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 408788394) // FIXME: should cache?
  var ret NodePath
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2DPhysicalBones) FetchPhysicalBones()  {
  classNameV := StringNameFromStr("SkeletonModification2DPhysicalBones")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("fetch_physical_bones")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DPhysicalBones) StartSimulation(bones StringName, )  {
  classNameV := StringNameFromStr("SkeletonModification2DPhysicalBones")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("start_simulation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2787316981) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bones.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DPhysicalBones) StopSimulation(bones StringName, )  {
  classNameV := StringNameFromStr("SkeletonModification2DPhysicalBones")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stop_simulation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2787316981) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bones.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

// Properties

func (me *SkeletonModification2DPhysicalBones) GetPropPhysicalBoneChainLength() int {
  panic("TODO: implement")
}

func (me *SkeletonModification2DPhysicalBones) SetPropPhysicalBoneChainLength(value int) {
  panic("TODO: implement")
}