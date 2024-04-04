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

type Shader struct {
  Resource
}

func (me *Shader) BaseClass() string {
  return "Shader"
}

func NewShader() *Shader {
  str := StringNameFromStr("Shader") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Shader{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

type ShaderMode int
const (
  ShaderModeModeSpatial ShaderMode = 0
  ShaderModeModeCanvasItem ShaderMode = 1
  ShaderModeModeParticles ShaderMode = 2
  ShaderModeModeSky ShaderMode = 3
  ShaderModeModeFog ShaderMode = 4
)

func (me *Shader) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Shader) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Shader) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Shader) GetMode() ShaderMode {
  classNameV := StringNameFromStr("Shader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_mode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3392948163) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret ShaderMode

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *Shader) SetCode(code String, )  {
  classNameV := StringNameFromStr("Shader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_code")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{code.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Shader) GetCode() String {
  classNameV := StringNameFromStr("Shader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_code")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Shader) SetDefaultTextureParameter(name StringName, texture Texture2D, index int64, )  {
  classNameV := StringNameFromStr("Shader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_default_texture_parameter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2750740428) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), texture.AsCTypePtr(), gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Shader) GetDefaultTextureParameter(name StringName, index int64, ) Texture2D {
  classNameV := StringNameFromStr("Shader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_default_texture_parameter")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3090538643) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{name.AsCTypePtr(), gdc.ConstTypePtr(&index) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewTexture2D()
  pinner.Pin(&index)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *Shader) GetShaderUniformList(get_groups bool, ) Array {
  classNameV := StringNameFromStr("Shader")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_shader_uniform_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1230511656) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&get_groups) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  pinner.Pin(&get_groups)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
