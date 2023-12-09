// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Shader struct {
  obj gdc.ObjectPtr
}

func (me *Shader) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Shader) BaseClass() string {
  return "Shader"
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

func (me *Shader) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Shader) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *Shader) GetMode()  {
  panic("TODO: implement")
}

func  (me *Shader) SetCode(code String, )  {
  panic("TODO: implement")
}

func  (me *Shader) GetCode()  {
  panic("TODO: implement")
}

func  (me *Shader) SetDefaultTextureParameter(name StringName, texture Texture2D, index int, )  {
  panic("TODO: implement")
}

func  (me *Shader) GetDefaultTextureParameter(name StringName, index int, )  {
  panic("TODO: implement")
}

func  (me *Shader) GetShaderUniformList(get_groups bool, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
