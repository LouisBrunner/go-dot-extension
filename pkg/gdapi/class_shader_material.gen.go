// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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
  panic("TODO: implement")
}

func  (me *ShaderMaterial) GetShader()  {
  panic("TODO: implement")
}

func  (me *ShaderMaterial) SetShaderParameter(param StringName, value Variant, )  {
  panic("TODO: implement")
}

func  (me *ShaderMaterial) GetShaderParameter(param StringName, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
