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

type VisualShaderNodeParameter struct {
  VisualShaderNode
}

func (me *VisualShaderNodeParameter) BaseClass() string {
  return "VisualShaderNodeParameter"
}

func NewVisualShaderNodeParameter() *VisualShaderNodeParameter {
  str := StringNameFromStr("VisualShaderNodeParameter") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeParameter{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type VisualShaderNodeParameterQualifier int
const (
  VisualShaderNodeParameterQualifierQualNone VisualShaderNodeParameterQualifier = 0
  VisualShaderNodeParameterQualifierQualGlobal VisualShaderNodeParameterQualifier = 1
  VisualShaderNodeParameterQualifierQualInstance VisualShaderNodeParameterQualifier = 2
  VisualShaderNodeParameterQualifierQualMax VisualShaderNodeParameterQualifier = 3
)

func (me *VisualShaderNodeParameter) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeParameter) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeParameter) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeParameter) SetParameterName(name String, )  {
  classNameV := StringNameFromStr("VisualShaderNodeParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_parameter_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeParameter) GetParameterName() String {
  classNameV := StringNameFromStr("VisualShaderNodeParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_parameter_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VisualShaderNodeParameter) SetQualifier(qualifier VisualShaderNodeParameterQualifier, )  {
  classNameV := StringNameFromStr("VisualShaderNodeParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_qualifier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1276489447) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&qualifier) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeParameter) GetQualifier() VisualShaderNodeParameterQualifier {
  classNameV := StringNameFromStr("VisualShaderNodeParameter")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_qualifier")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3558406205) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeParameterQualifier

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
