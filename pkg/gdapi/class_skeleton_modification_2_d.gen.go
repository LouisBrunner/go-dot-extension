// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SkeletonModification2D struct {
  Resource
}

func (me *SkeletonModification2D) BaseClass() string {
  return "SkeletonModification2D"
}

func NewSkeletonModification2D() *SkeletonModification2D {
  str := StringNameFromStr("SkeletonModification2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SkeletonModification2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *SkeletonModification2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SkeletonModification2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SkeletonModification2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SkeletonModification2D) SetEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("SkeletonModification2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2D) GetEnabled() bool {
  classNameV := StringNameFromStr("SkeletonModification2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2D) GetModificationStack() SkeletonModificationStack2D {
  classNameV := StringNameFromStr("SkeletonModification2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_modification_stack")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2137761694) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewSkeletonModificationStack2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SkeletonModification2D) SetIsSetup(is_setup bool, )  {
  classNameV := StringNameFromStr("SkeletonModification2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_is_setup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&is_setup), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2D) GetIsSetup() bool {
  classNameV := StringNameFromStr("SkeletonModification2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_is_setup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2D) SetExecutionMode(execution_mode int64, )  {
  classNameV := StringNameFromStr("SkeletonModification2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_execution_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&execution_mode), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2D) GetExecutionMode() int64 {
  classNameV := StringNameFromStr("SkeletonModification2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_execution_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2D) ClampAngle(angle float64, min float64, max float64, invert bool, ) float64 {
  classNameV := StringNameFromStr("SkeletonModification2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clamp_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1229502682) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angle), gdc.ConstTypePtr(&min), gdc.ConstTypePtr(&max), gdc.ConstTypePtr(&invert), }
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2D) SetEditorDrawGizmo(draw_gizmo bool, )  {
  classNameV := StringNameFromStr("SkeletonModification2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_editor_draw_gizmo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draw_gizmo), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2D) GetEditorDrawGizmo() bool {
  classNameV := StringNameFromStr("SkeletonModification2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_editor_draw_gizmo")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
