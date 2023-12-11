// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ShaderMaterial struct {
  obj gdc.ObjectPtr
}

func (me *ShaderMaterial) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ShaderMaterial) BaseClass() string {
  return "ShaderMaterial"
}



// Enums

func (me *ShaderMaterial) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ShaderMaterial) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ShaderMaterial) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ShaderMaterial) SetShader(shader Shader, )  {
  classNameV := StringNameFromStr("ShaderMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shader")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341921675) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(shader.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShaderMaterial) GetShader() Shader {
  classNameV := StringNameFromStr("ShaderMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shader")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2078273437) // FIXME: should cache?
  var ret Shader
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *ShaderMaterial) SetShaderParameter(param StringName, value Variant, )  {
  classNameV := StringNameFromStr("ShaderMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shader_parameter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3776071444) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(param.AsCTypePtr()), gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ShaderMaterial) GetShaderParameter(param StringName, ) Variant {
  classNameV := StringNameFromStr("ShaderMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shader_parameter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2760726917) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(param.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
