// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SkeletonModificationStack2D struct {
  obj gdc.ObjectPtr
}

func (me *SkeletonModificationStack2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SkeletonModificationStack2D) BaseClass() string {
  return "SkeletonModificationStack2D"
}



// Enums

func (me *SkeletonModificationStack2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SkeletonModificationStack2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SkeletonModificationStack2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SkeletonModificationStack2D) Setup()  {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("setup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3218959716) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModificationStack2D) Execute(delta float32, execution_mode int, )  {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("execute")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1005356550) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delta), gdc.ConstTypePtr(&execution_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModificationStack2D) EnableAllModifications(enabled bool, )  {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("enable_all_modifications")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModificationStack2D) GetModification(mod_idx int, ) SkeletonModification2D {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_modification")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2570274329) // FIXME: should cache?
  var ret SkeletonModification2D
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mod_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModificationStack2D) AddModification(modification SkeletonModification2D, )  {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_modification")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 354162120) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(modification.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModificationStack2D) DeleteModification(mod_idx int, )  {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("delete_modification")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mod_idx), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModificationStack2D) SetModification(mod_idx int, modification SkeletonModification2D, )  {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_modification")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1098262544) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mod_idx), gdc.ConstTypePtr(modification.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModificationStack2D) SetModificationCount(count int, )  {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_modification_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModificationStack2D) GetModificationCount() int {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_modification_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModificationStack2D) GetIsSetup() bool {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_is_setup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModificationStack2D) SetEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModificationStack2D) GetEnabled() bool {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModificationStack2D) SetStrength(strength float32, )  {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&strength), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModificationStack2D) GetStrength() float32 {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModificationStack2D) GetSkeleton() Skeleton2D {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skeleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1697361217) // FIXME: should cache?
  var ret Skeleton2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
