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

type VisualShaderNodeVec4Parameter struct {
  VisualShaderNodeParameter
}

func (me *VisualShaderNodeVec4Parameter) BaseClass() string {
  return "VisualShaderNodeVec4Parameter"
}

func NewVisualShaderNodeVec4Parameter() *VisualShaderNodeVec4Parameter {
  str := StringNameFromStr("VisualShaderNodeVec4Parameter") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeVec4Parameter{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeVec4Parameter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVec4Parameter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVec4Parameter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeVec4Parameter) SetDefaultValueEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("VisualShaderNodeVec4Parameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_value_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeVec4Parameter) IsDefaultValueEnabled() bool {
  classNameV := StringNameFromStr("VisualShaderNodeVec4Parameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_default_value_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *VisualShaderNodeVec4Parameter) SetDefaultValue(value Vector4, )  {
  classNameV := StringNameFromStr("VisualShaderNodeVec4Parameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 643568085) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeVec4Parameter) GetDefaultValue() Vector4 {
  classNameV := StringNameFromStr("VisualShaderNodeVec4Parameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_value")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2435802345) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector4()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
