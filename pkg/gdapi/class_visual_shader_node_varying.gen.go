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

type VisualShaderNodeVarying struct {
  VisualShaderNode
}

func (me *VisualShaderNodeVarying) BaseClass() string {
  return "VisualShaderNodeVarying"
}

func NewVisualShaderNodeVarying() *VisualShaderNodeVarying {
  str := StringNameFromStr("VisualShaderNodeVarying") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeVarying{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *VisualShaderNodeVarying) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeVarying) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeVarying) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeVarying) SetVaryingName(name String, )  {
  classNameV := StringNameFromStr("VisualShaderNodeVarying")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_varying_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeVarying) GetVaryingName() String {
  classNameV := StringNameFromStr("VisualShaderNodeVarying")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_varying_name")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *VisualShaderNodeVarying) SetVaryingType(type_ VisualShaderVaryingType, )  {
  classNameV := StringNameFromStr("VisualShaderNodeVarying")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_varying_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3565867981) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&type_) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeVarying) GetVaryingType() VisualShaderVaryingType {
  classNameV := StringNameFromStr("VisualShaderNodeVarying")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_varying_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 523183580) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderVaryingType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
