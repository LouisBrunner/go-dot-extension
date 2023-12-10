// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeInput struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeInput) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
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

func (me *VisualShaderNodeInput) GetPropInputName() StringName {
  panic("TODO: implement")
}

func (me *VisualShaderNodeInput) SetPropInputName(value StringName) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API