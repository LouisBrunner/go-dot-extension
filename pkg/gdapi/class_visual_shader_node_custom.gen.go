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

type ptrsForVisualShaderNodeCustomList struct {
  fnXGetName gdc.MethodBindPtr
  fnXGetDescription gdc.MethodBindPtr
  fnXGetCategory gdc.MethodBindPtr
  fnXGetReturnIconType gdc.MethodBindPtr
  fnXGetInputPortCount gdc.MethodBindPtr
  fnXGetInputPortType gdc.MethodBindPtr
  fnXGetInputPortName gdc.MethodBindPtr
  fnXGetInputPortDefaultValue gdc.MethodBindPtr
  fnXGetDefaultInputPort gdc.MethodBindPtr
  fnXGetOutputPortCount gdc.MethodBindPtr
  fnXGetOutputPortType gdc.MethodBindPtr
  fnXGetOutputPortName gdc.MethodBindPtr
  fnXGetPropertyCount gdc.MethodBindPtr
  fnXGetPropertyName gdc.MethodBindPtr
  fnXGetPropertyDefaultIndex gdc.MethodBindPtr
  fnXGetPropertyOptions gdc.MethodBindPtr
  fnXGetCode gdc.MethodBindPtr
  fnXGetFuncCode gdc.MethodBindPtr
  fnXGetGlobalCode gdc.MethodBindPtr
  fnXIsHighend gdc.MethodBindPtr
  fnXIsAvailable gdc.MethodBindPtr
  fnGetOptionIndex gdc.MethodBindPtr
}

var ptrsForVisualShaderNodeCustom ptrsForVisualShaderNodeCustomList

func initVisualShaderNodeCustomPtrs(iface gdc.Interface) {

  className := StringNameFromStr("VisualShaderNodeCustom")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("get_option_index")
    defer methodName.Destroy()
    ptrsForVisualShaderNodeCustom.fnGetOptionIndex = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 923996154))
  }
}

type VisualShaderNodeCustom struct {
  VisualShaderNode
}

func (me *VisualShaderNodeCustom) BaseClass() string {
  return "VisualShaderNodeCustom"
}

func NewVisualShaderNodeCustom() *VisualShaderNodeCustom {
  str := StringNameFromStr("VisualShaderNodeCustom") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeCustom{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeCustom) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeCustom) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeCustom) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeCustom) GetOptionIndex(option int64, ) int64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&option) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()
  pinner.Pin(&option)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForVisualShaderNodeCustom.fnGetOptionIndex), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
