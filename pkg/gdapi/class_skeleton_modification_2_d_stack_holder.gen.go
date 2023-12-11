// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SkeletonModification2DStackHolder struct {
  obj gdc.ObjectPtr
}

func (me *SkeletonModification2DStackHolder) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SkeletonModification2DStackHolder) BaseClass() string {
  return "SkeletonModification2DStackHolder"
}



// Enums

func (me *SkeletonModification2DStackHolder) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SkeletonModification2DStackHolder) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SkeletonModification2DStackHolder) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SkeletonModification2DStackHolder) SetHeldModificationStack(held_modification_stack SkeletonModificationStack2D, )  {
  classNameV := StringNameFromStr("SkeletonModification2DStackHolder")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_held_modification_stack")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3907307132) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(held_modification_stack.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2DStackHolder) GetHeldModificationStack() SkeletonModificationStack2D {
  classNameV := StringNameFromStr("SkeletonModification2DStackHolder")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_held_modification_stack")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2107508396) // FIXME: should cache?
  var ret SkeletonModificationStack2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
