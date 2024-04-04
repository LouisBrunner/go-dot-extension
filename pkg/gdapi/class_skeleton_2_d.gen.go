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

type Skeleton2D struct {
  Node2D
}

func (me *Skeleton2D) BaseClass() string {
  return "Skeleton2D"
}

func NewSkeleton2D() *Skeleton2D {
  str := StringNameFromStr("Skeleton2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Skeleton2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *Skeleton2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Skeleton2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Skeleton2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Skeleton2D) GetBoneCount() int64 {
  classNameV := StringNameFromStr("Skeleton2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Skeleton2D) GetBone(idx int64, ) Bone2D {
  classNameV := StringNameFromStr("Skeleton2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2556267111) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBone2D()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skeleton2D) GetSkeleton() RID {
  classNameV := StringNameFromStr("Skeleton2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skeleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skeleton2D) SetModificationStack(modification_stack SkeletonModificationStack2D, )  {
  classNameV := StringNameFromStr("Skeleton2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_modification_stack")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3907307132) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{modification_stack.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton2D) GetModificationStack() SkeletonModificationStack2D {
  classNameV := StringNameFromStr("Skeleton2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_modification_stack")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2107508396) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSkeletonModificationStack2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skeleton2D) ExecuteModifications(delta float64, execution_mode int64, )  {
  classNameV := StringNameFromStr("Skeleton2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("execute_modifications")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1005356550) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delta) , gdc.ConstTypePtr(&execution_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton2D) SetBoneLocalPoseOverride(bone_idx int64, override_pose Transform2D, strength float64, persistent bool, )  {
  classNameV := StringNameFromStr("Skeleton2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone_local_pose_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 555457532) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , override_pose.AsCTypePtr(), gdc.ConstTypePtr(&strength) , gdc.ConstTypePtr(&persistent) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton2D) GetBoneLocalPoseOverride(bone_idx int64, ) Transform2D {
  classNameV := StringNameFromStr("Skeleton2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_local_pose_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2995540667) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()
  pinner.Pin(&bone_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals

type Skeleton2DBoneSetupChangedSignalFn func()

func (me *Skeleton2D) ConnectBoneSetupChanged(subs SignalSubscribers, fn Skeleton2DBoneSetupChangedSignalFn) {
  sig := StringNameFromStr("bone_setup_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *Skeleton2D) DisconnectBoneSetupChanged(subs SignalSubscribers, fn Skeleton2DBoneSetupChangedSignalFn) {
  sig := StringNameFromStr("bone_setup_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
