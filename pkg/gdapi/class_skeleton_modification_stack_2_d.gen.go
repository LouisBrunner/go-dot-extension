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

type SkeletonModificationStack2D struct {
  Resource
}

func (me *SkeletonModificationStack2D) BaseClass() string {
  return "SkeletonModificationStack2D"
}

func NewSkeletonModificationStack2D() *SkeletonModificationStack2D {
  str := StringNameFromStr("SkeletonModificationStack2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SkeletonModificationStack2D{}
  obj.SetBaseObject(objPtr)
  return obj
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
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModificationStack2D) Execute(delta float64, execution_mode int64, )  {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("execute")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1005356550) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&delta) , gdc.ConstTypePtr(&execution_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModificationStack2D) EnableAllModifications(enabled bool, )  {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("enable_all_modifications")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModificationStack2D) GetModification(mod_idx int64, ) SkeletonModification2D {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_modification")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2570274329) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mod_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSkeletonModification2D()
  pinner.Pin(&mod_idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SkeletonModificationStack2D) AddModification(modification SkeletonModification2D, )  {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("add_modification")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 354162120) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{modification.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModificationStack2D) DeleteModification(mod_idx int64, )  {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("delete_modification")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mod_idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModificationStack2D) SetModification(mod_idx int64, modification SkeletonModification2D, )  {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_modification")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1098262544) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mod_idx) , modification.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModificationStack2D) SetModificationCount(count int64, )  {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_modification_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&count) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModificationStack2D) GetModificationCount() int64 {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_modification_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModificationStack2D) GetIsSetup() bool {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_is_setup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModificationStack2D) SetEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModificationStack2D) GetEnabled() bool {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModificationStack2D) SetStrength(strength float64, )  {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&strength) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModificationStack2D) GetStrength() float64 {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_strength")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModificationStack2D) GetSkeleton() Skeleton2D {
  classNameV := StringNameFromStr("SkeletonModificationStack2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_skeleton")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1697361217) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSkeleton2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
