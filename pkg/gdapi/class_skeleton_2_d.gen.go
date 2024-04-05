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

type ptrsForSkeleton2DList struct {
  fnGetBoneCount gdc.MethodBindPtr
  fnGetBone gdc.MethodBindPtr
  fnGetSkeleton gdc.MethodBindPtr
  fnSetModificationStack gdc.MethodBindPtr
  fnGetModificationStack gdc.MethodBindPtr
  fnExecuteModifications gdc.MethodBindPtr
  fnSetBoneLocalPoseOverride gdc.MethodBindPtr
  fnGetBoneLocalPoseOverride gdc.MethodBindPtr
}

var ptrsForSkeleton2D ptrsForSkeleton2DList

func initSkeleton2DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Skeleton2D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_bone_count")
    defer methodName.Destroy()
    ptrsForSkeleton2D.fnGetBoneCount = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("get_bone")
    defer methodName.Destroy()
    ptrsForSkeleton2D.fnGetBone = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2556267111))
  }
  {
    methodName := StringNameFromStr("get_skeleton")
    defer methodName.Destroy()
    ptrsForSkeleton2D.fnGetSkeleton = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2944877500))
  }
  {
    methodName := StringNameFromStr("set_modification_stack")
    defer methodName.Destroy()
    ptrsForSkeleton2D.fnSetModificationStack = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3907307132))
  }
  {
    methodName := StringNameFromStr("get_modification_stack")
    defer methodName.Destroy()
    ptrsForSkeleton2D.fnGetModificationStack = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2107508396))
  }
  {
    methodName := StringNameFromStr("execute_modifications")
    defer methodName.Destroy()
    ptrsForSkeleton2D.fnExecuteModifications = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1005356550))
  }
  {
    methodName := StringNameFromStr("set_bone_local_pose_override")
    defer methodName.Destroy()
    ptrsForSkeleton2D.fnSetBoneLocalPoseOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 555457532))
  }
  {
    methodName := StringNameFromStr("get_bone_local_pose_override")
    defer methodName.Destroy()
    ptrsForSkeleton2D.fnGetBoneLocalPoseOverride = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2995540667))
  }
}

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
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton2D.fnGetBoneCount), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *Skeleton2D) GetBone(idx int64, ) Bone2D {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBone2D()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton2D.fnGetBone), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skeleton2D) GetSkeleton() RID {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRID()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton2D.fnGetSkeleton), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skeleton2D) SetModificationStack(modification_stack SkeletonModificationStack2D, )  {
  cargs := []gdc.ConstTypePtr{modification_stack.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton2D.fnSetModificationStack), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton2D) GetModificationStack() SkeletonModificationStack2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSkeletonModificationStack2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton2D.fnGetModificationStack), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Skeleton2D) ExecuteModifications(delta float64, execution_mode int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delta) , gdc.ConstTypePtr(&execution_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton2D.fnExecuteModifications), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton2D) SetBoneLocalPoseOverride(bone_idx int64, override_pose Transform2D, strength float64, persistent bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , override_pose.AsCTypePtr(), gdc.ConstTypePtr(&strength) , gdc.ConstTypePtr(&persistent) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton2D.fnSetBoneLocalPoseOverride), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Skeleton2D) GetBoneLocalPoseOverride(bone_idx int64, ) Transform2D {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&bone_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTransform2D()
  pinner.Pin(&bone_idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeleton2D.fnGetBoneLocalPoseOverride), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
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
