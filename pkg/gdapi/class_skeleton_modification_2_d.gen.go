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

type ptrsForSkeletonModification2DList struct {
  fnXExecute gdc.MethodBindPtr
  fnXSetupModification gdc.MethodBindPtr
  fnXDrawEditorGizmo gdc.MethodBindPtr
  fnSetEnabled gdc.MethodBindPtr
  fnGetEnabled gdc.MethodBindPtr
  fnGetModificationStack gdc.MethodBindPtr
  fnSetIsSetup gdc.MethodBindPtr
  fnGetIsSetup gdc.MethodBindPtr
  fnSetExecutionMode gdc.MethodBindPtr
  fnGetExecutionMode gdc.MethodBindPtr
  fnClampAngle gdc.MethodBindPtr
  fnSetEditorDrawGizmo gdc.MethodBindPtr
  fnGetEditorDrawGizmo gdc.MethodBindPtr
}

var ptrsForSkeletonModification2D ptrsForSkeletonModification2DList

func initSkeletonModification2DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("SkeletonModification2D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_enabled")
    defer methodName.Destroy()
    ptrsForSkeletonModification2D.fnSetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_enabled")
    defer methodName.Destroy()
    ptrsForSkeletonModification2D.fnGetEnabled = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2240911060))
  }
  {
    methodName := StringNameFromStr("get_modification_stack")
    defer methodName.Destroy()
    ptrsForSkeletonModification2D.fnGetModificationStack = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2137761694))
  }
  {
    methodName := StringNameFromStr("set_is_setup")
    defer methodName.Destroy()
    ptrsForSkeletonModification2D.fnSetIsSetup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_is_setup")
    defer methodName.Destroy()
    ptrsForSkeletonModification2D.fnGetIsSetup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
  {
    methodName := StringNameFromStr("set_execution_mode")
    defer methodName.Destroy()
    ptrsForSkeletonModification2D.fnSetExecutionMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1286410249))
  }
  {
    methodName := StringNameFromStr("get_execution_mode")
    defer methodName.Destroy()
    ptrsForSkeletonModification2D.fnGetExecutionMode = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3905245786))
  }
  {
    methodName := StringNameFromStr("clamp_angle")
    defer methodName.Destroy()
    ptrsForSkeletonModification2D.fnClampAngle = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1229502682))
  }
  {
    methodName := StringNameFromStr("set_editor_draw_gizmo")
    defer methodName.Destroy()
    ptrsForSkeletonModification2D.fnSetEditorDrawGizmo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2586408642))
  }
  {
    methodName := StringNameFromStr("get_editor_draw_gizmo")
    defer methodName.Destroy()
    ptrsForSkeletonModification2D.fnGetEditorDrawGizmo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 36873697))
  }
}

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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2D.fnSetEnabled), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2D) GetEnabled() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2D.fnGetEnabled), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2D) GetModificationStack() SkeletonModificationStack2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewSkeletonModificationStack2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2D.fnGetModificationStack), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *SkeletonModification2D) SetIsSetup(is_setup bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&is_setup) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2D.fnSetIsSetup), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2D) GetIsSetup() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2D.fnGetIsSetup), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2D) SetExecutionMode(execution_mode int64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&execution_mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2D.fnSetExecutionMode), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2D) GetExecutionMode() int64 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2D.fnGetExecutionMode), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2D) ClampAngle(angle float64, min float64, max float64, invert bool, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&angle) , gdc.ConstTypePtr(&min) , gdc.ConstTypePtr(&max) , gdc.ConstTypePtr(&invert) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&angle)
  pinner.Pin(&min)
  pinner.Pin(&max)
  pinner.Pin(&invert)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2D.fnClampAngle), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SkeletonModification2D) SetEditorDrawGizmo(draw_gizmo bool, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&draw_gizmo) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2D.fnSetEditorDrawGizmo), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SkeletonModification2D) GetEditorDrawGizmo() bool {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForSkeletonModification2D.fnGetEditorDrawGizmo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
