// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeInput struct {
  VisualShaderNode
}

func (me *VisualShaderNodeInput) BaseClass() string {
  return "VisualShaderNodeInput"
}



// Enums

func (me *VisualShaderNodeInput) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeInput) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeInput) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeInput) SetInputName(name String, )  {
  classNameV := StringNameFromStr("VisualShaderNodeInput")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_input_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(name.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeInput) GetInputName() String {
  classNameV := StringNameFromStr("VisualShaderNodeInput")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeInput) GetInputRealName() String {
  classNameV := StringNameFromStr("VisualShaderNodeInput")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_input_real_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals

type VisualShaderNodeInputInputTypeChangedSignalFn func()

func (me *VisualShaderNodeInput) ConnectInputTypeChanged(subs SignalSubscribers, fn VisualShaderNodeInputInputTypeChangedSignalFn) {
  sig := StringNameFromStr("input_type_changed")
  defer sig.Destroy()
  me.Connect(*sig, subs.add(fn), 0)
}

func (me *VisualShaderNodeInput) DisconnectInputTypeChanged(subs SignalSubscribers, fn VisualShaderNodeInputInputTypeChangedSignalFn) {
  sig := StringNameFromStr("input_type_changed")
  defer sig.Destroy()
  me.Disconnect(*sig, *subs.remove(fn))
}
