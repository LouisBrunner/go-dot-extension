// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type AnimationNodeStateMachineTransition struct {
  obj gdc.ObjectPtr
}

func (me *AnimationNodeStateMachineTransition) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimationNodeStateMachineTransition) BaseClass() string {
  return "AnimationNodeStateMachineTransition"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeStateMachineTransition) GetSwitchMode() AnimationNodeStateMachineTransitionSwitchMode {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_switch_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2138562085) // FIXME: should cache?
  var ret AnimationNodeStateMachineTransitionSwitchMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeStateMachineTransition) SetAdvanceMode(mode AnimationNodeStateMachineTransitionAdvanceMode, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_advance_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1210869868) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&mode), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeStateMachineTransition) GetAdvanceMode() AnimationNodeStateMachineTransitionAdvanceMode {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_advance_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 61101689) // FIXME: should cache?
  var ret AnimationNodeStateMachineTransitionAdvanceMode
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeStateMachineTransition) SetAdvanceCondition(name StringName, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_advance_condition")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3304788590) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeStateMachineTransition) GetAdvanceCondition() StringName {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_advance_condition")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2002593661) // FIXME: should cache?
  var ret StringName
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeStateMachineTransition) SetXfadeTime(secs float32, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_xfade_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&secs), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeStateMachineTransition) GetXfadeTime() float32 {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_xfade_time")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeStateMachineTransition) SetXfadeCurve(curve Curve, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_xfade_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 270443179) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(curve.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeStateMachineTransition) GetXfadeCurve() Curve {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_xfade_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2460114913) // FIXME: should cache?
  var ret Curve
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeStateMachineTransition) SetReset(reset bool, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_reset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&reset), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeStateMachineTransition) IsReset() bool {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_reset")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeStateMachineTransition) SetPriority(priority int, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&priority), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeStateMachineTransition) GetPriority() int {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_priority")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *AnimationNodeStateMachineTransition) SetAdvanceExpression(text String, )  {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_advance_expression")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(text.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *AnimationNodeStateMachineTransition) GetAdvanceExpression() String {
  classNameV := StringNameFromStr("AnimationNodeStateMachineTransition")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_advance_expression")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *AnimationNodeStateMachineTransition) GetPropXfadeTime() float32 {
  panic("TODO: implement")
}

func (me *AnimationNodeStateMachineTransition) SetPropXfadeTime(value float32) {
  panic("TODO: implement")
}

func (me *AnimationNodeStateMachineTransition) GetPropXfadeCurve() Curve {
  panic("TODO: implement")
}

func (me *AnimationNodeStateMachineTransition) SetPropXfadeCurve(value Curve) {
  panic("TODO: implement")
}

func (me *AnimationNodeStateMachineTransition) GetPropReset() bool {
  panic("TODO: implement")
}

func (me *AnimationNodeStateMachineTransition) SetPropReset(value bool) {
  panic("TODO: implement")
}

func (me *AnimationNodeStateMachineTransition) GetPropPriority() int {
  panic("TODO: implement")
}

func (me *AnimationNodeStateMachineTransition) SetPropPriority(value int) {
  panic("TODO: implement")
}

func (me *AnimationNodeStateMachineTransition) GetPropSwitchMode() int {
  panic("TODO: implement")
}

func (me *AnimationNodeStateMachineTransition) SetPropSwitchMode(value int) {
  panic("TODO: implement")
}

func (me *AnimationNodeStateMachineTransition) GetPropAdvanceMode() int {
  panic("TODO: implement")
}

func (me *AnimationNodeStateMachineTransition) SetPropAdvanceMode(value int) {
  panic("TODO: implement")
}

func (me *AnimationNodeStateMachineTransition) GetPropAdvanceCondition() StringName {
  panic("TODO: implement")
}

func (me *AnimationNodeStateMachineTransition) SetPropAdvanceCondition(value StringName) {
  panic("TODO: implement")
}

func (me *AnimationNodeStateMachineTransition) GetPropAdvanceExpression() String {
  panic("TODO: implement")
}

func (me *AnimationNodeStateMachineTransition) SetPropAdvanceExpression(value String) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API