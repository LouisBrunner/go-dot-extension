// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Skeleton2D struct {
  Node2D
}

func (me *Skeleton2D) BaseClass() string {
  return "Skeleton2D"
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

func  (me *Skeleton2D) GetBoneCount() int {
  classNameV := StringNameFromStr("Skeleton2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Skeleton2D) GetBone(idx int, ) Bone2D {
  classNameV := StringNameFromStr("Skeleton2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2556267111) // FIXME: should cache?
  var ret Bone2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Skeleton2D) GetSkeleton() RID {
  classNameV := StringNameFromStr("Skeleton2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skeleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2944877500) // FIXME: should cache?
  var ret RID
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Skeleton2D) SetModificationStack(modification_stack SkeletonModificationStack2D, )  {
  classNameV := StringNameFromStr("Skeleton2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_modification_stack")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3907307132) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(modification_stack.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Skeleton2D) GetModificationStack() SkeletonModificationStack2D {
  classNameV := StringNameFromStr("Skeleton2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_modification_stack")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2107508396) // FIXME: should cache?
  var ret SkeletonModificationStack2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *Skeleton2D) ExecuteModifications(delta float32, execution_mode int, )  {
  classNameV := StringNameFromStr("Skeleton2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("execute_modifications")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1005356550) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delta), gdc.ConstTypePtr(&execution_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Skeleton2D) SetBoneLocalPoseOverride(bone_idx int, override_pose Transform2D, strength float32, persistent bool, )  {
  classNameV := StringNameFromStr("Skeleton2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bone_local_pose_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 555457532) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), gdc.ConstTypePtr(override_pose.AsCTypePtr()), gdc.ConstTypePtr(&strength), gdc.ConstTypePtr(&persistent), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Skeleton2D) GetBoneLocalPoseOverride(bone_idx int, ) Transform2D {
  classNameV := StringNameFromStr("Skeleton2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bone_local_pose_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2995540667) // FIXME: should cache?
  var ret Transform2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals

type Skeleton2DBoneSetupChangedSignalFn func()

func (me *Skeleton2D) ConnectBoneSetupChanged(subs SignalSubscribers, fn Skeleton2DBoneSetupChangedSignalFn) {
  sig := StringNameFromStr("bone_setup_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Connect(sig, subs.add(fn), 0)
}

func (me *Skeleton2D) DisconnectBoneSetupChanged(subs SignalSubscribers, fn Skeleton2DBoneSetupChangedSignalFn) {
  sig := StringNameFromStr("bone_setup_changed")
  defer sig.Destroy()
  obj := ObjectFromPtr(me.obj)
  obj.Disconnect(sig, *subs.remove(fn))
}
