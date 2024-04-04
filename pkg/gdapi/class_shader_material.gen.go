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

type ShaderMaterial struct {
  Material
}

func (me *ShaderMaterial) BaseClass() string {
  return "ShaderMaterial"
}

func NewShaderMaterial() *ShaderMaterial {
  str := StringNameFromStr("ShaderMaterial") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ShaderMaterial{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{shader.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShaderMaterial) GetShader() Shader {
  classNameV := StringNameFromStr("ShaderMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shader")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2078273437) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewShader()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *ShaderMaterial) SetShaderParameter(param StringName, value Variant, )  {
  classNameV := StringNameFromStr("ShaderMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_shader_parameter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3776071444) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{param.AsCTypePtr(), value.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ShaderMaterial) GetShaderParameter(param StringName, ) Variant {
  classNameV := StringNameFromStr("ShaderMaterial")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shader_parameter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2760726917) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{param.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVariant()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
