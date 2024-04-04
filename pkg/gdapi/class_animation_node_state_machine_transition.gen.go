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

type AnimationNodeStateMachineTransition struct {
  Resource
}

func (me *AnimationNodeStateMachineTransition) BaseClass() string {
  return "AnimationNodeStateMachineTransition"
}

func NewAnimationNodeStateMachineTransition() *AnimationNodeStateMachineTransition {
  str := StringNameFromStr("AnimationNodeStateMachineTransition") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &AnimationNodeStateMachineTransition{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type AnimationNodeStateMachineTransitionSwitchMode int
const (
  AnimationNodeStateMachineTransitionSwitchModeSwitchModeImmediate AnimationNodeStateMachineTransitionSwitchMode = 0
  AnimationNodeStateMachineTransitionSwitchModeSwitchModeSync AnimationNodeStateMachineTransitionSwitchMode = 1
  AnimationNodeStateMachineTransitionSwitchModeSwitchModeAtEnd AnimationNodeStateMachineTransitionSwitchMode = 2
)

type AnimationNodeStateMachineTransitionAdvanceMode int
const (
  AnimationNodeStateMachineTransitionAdvanceModeAdvanceModeDisabled AnimationNodeStateMachineTransitionAdvanceMode = 0
  AnimationNodeStateMachineTransitionAdvanceModeAdvanceModeEnabled AnimationNodeStateMachineTransitionAdvanceMode = 1
  AnimationNodeStateMachineTransitionAdvanceModeAdvanceModeAuto AnimationNodeStateMachineTransitionAdvanceMode = 2
)

func (me *AnimationNodeStateMachineTransition) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimationNodeStateMachineTransition) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimationNodeStateMachineTransition) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *AnimationNodeStateMachineTransition) SetSwitchMode(mode AnimationNodeStateMachineTransitionSwitchMode, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_switch_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2074906633) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachineTransition) GetSwitchMode() AnimationNodeStateMachineTransitionSwitchMode {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_switch_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2138562085) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AnimationNodeStateMachineTransitionSwitchMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AnimationNodeStateMachineTransition) SetAdvanceMode(mode AnimationNodeStateMachineTransitionAdvanceMode, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_advance_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1210869868) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachineTransition) GetAdvanceMode() AnimationNodeStateMachineTransitionAdvanceMode {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_advance_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 61101689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret AnimationNodeStateMachineTransitionAdvanceMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *AnimationNodeStateMachineTransition) SetAdvanceCondition(name StringName, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_advance_condition")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachineTransition) GetAdvanceCondition() StringName {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_advance_condition")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewStringName()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeStateMachineTransition) SetXfadeTime(secs float64, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_xfade_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachineTransition) GetXfadeTime() float64 {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_xfade_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeStateMachineTransition) SetXfadeCurve(curve Curve, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_xfade_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 270443179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{curve.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachineTransition) GetXfadeCurve() Curve {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_xfade_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2460114913) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCurve()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *AnimationNodeStateMachineTransition) SetReset(reset bool, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_reset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&reset) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachineTransition) IsReset() bool {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_reset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeStateMachineTransition) SetPriority(priority int64, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachineTransition) GetPriority() int64 {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *AnimationNodeStateMachineTransition) SetAdvanceExpression(text String, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_advance_expression")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{text.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *AnimationNodeStateMachineTransition) GetAdvanceExpression() String {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_advance_expression")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type AnimationNodeStateMachineTransitionAdvanceConditionChangedSignalFn func()

func (me *AnimationNodeStateMachineTransition) ConnectAdvanceConditionChanged(subs SignalSubscribers, fn AnimationNodeStateMachineTransitionAdvanceConditionChangedSignalFn) {
  sig := StringNameFromStr("advance_condition_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *AnimationNodeStateMachineTransition) DisconnectAdvanceConditionChanged(subs SignalSubscribers, fn AnimationNodeStateMachineTransitionAdvanceConditionChangedSignalFn) {
  sig := StringNameFromStr("advance_condition_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
