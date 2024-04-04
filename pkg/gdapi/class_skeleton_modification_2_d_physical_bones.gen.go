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

type SkeletonModification2DPhysicalBones struct {
  SkeletonModification2D
}

func (me *SkeletonModification2DPhysicalBones) BaseClass() string {
  return "SkeletonModification2DPhysicalBones"
}

func NewSkeletonModification2DPhysicalBones() *SkeletonModification2DPhysicalBones {
  str := StringNameFromStr("SkeletonModification2DPhysicalBones") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SkeletonModification2DPhysicalBones{}
  obj.SetBaseObject(objPtr)
  return obj
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

func  (me *SkeletonModification2DPhysicalBones) SetPhysicalBoneChainLength(length int64, )  {
  classNameV := StringNameFromStr("SkeletonModification2DPhysicalBones")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_physical_bone_chain_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&length) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DPhysicalBones) GetPhysicalBoneChainLength() int64 {
  classNameV := StringNameFromStr("SkeletonModification2DPhysicalBones")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_physical_bone_chain_length")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2455072627) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2DPhysicalBones) SetPhysicalBoneNode(joint_idx int64, physicalbone2d_node NodePath, )  {
  classNameV := StringNameFromStr("SkeletonModification2DPhysicalBones")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_physical_bone_node")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2761262315) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&joint_idx) , physicalbone2d_node.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DPhysicalBones) GetPhysicalBoneNode(joint_idx int64, ) NodePath {
  classNameV := StringNameFromStr("SkeletonModification2DPhysicalBones")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_physical_bone_node")
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

func  (me *SkeletonModification2DPhysicalBones) FetchPhysicalBones()  {
  classNameV := StringNameFromStr("SkeletonModification2DPhysicalBones")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("fetch_physical_bones")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DPhysicalBones) StartSimulation(bones []StringName, )  {
  classNameV := StringNameFromStr("SkeletonModification2DPhysicalBones")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("start_simulation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2787316981) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bones) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2DPhysicalBones) StopSimulation(bones []StringName, )  {
  classNameV := StringNameFromStr("SkeletonModification2DPhysicalBones")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("stop_simulation")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2787316981) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bones) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
