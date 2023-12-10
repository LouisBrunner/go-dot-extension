// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SkeletonModification2D struct {
  obj gdc.ObjectPtr
}

func (me *SkeletonModification2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SkeletonModification2D) BaseClass() string {
  return "SkeletonModification2D"
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
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2D) GetModificationStack() SkeletonModificationStack2D {
  classNameV := StringNameFromStr("SkeletonModification2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_modification_stack")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2137761694) // FIXME: should cache?
  var ret SkeletonModificationStack2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2D) SetExecutionMode(execution_mode int, )  {
  classNameV := StringNameFromStr("SkeletonModification2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_execution_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&execution_mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SkeletonModification2D) GetExecutionMode() int {
  classNameV := StringNameFromStr("SkeletonModification2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_execution_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SkeletonModification2D) ClampAngle(angle float32, min float32, max float32, invert bool, ) float32 {
  classNameV := StringNameFromStr("SkeletonModification2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("clamp_angle")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1229502682) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angle), gdc.ConstTypePtr(&min), gdc.ConstTypePtr(&max), gdc.ConstTypePtr(&invert), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
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
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *SkeletonModification2D) GetPropEnabled() bool {
  panic("TODO: implement")
}

func (me *SkeletonModification2D) SetPropEnabled(value bool) {
  panic("TODO: implement")
}

func (me *SkeletonModification2D) GetPropExecutionMode() int {
  panic("TODO: implement")
}

func (me *SkeletonModification2D) SetPropExecutionMode(value int) {
  panic("TODO: implement")
}