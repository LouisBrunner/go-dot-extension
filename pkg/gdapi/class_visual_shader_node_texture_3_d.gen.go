// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTexture3D struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeTexture3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeTexture3D) BaseClass() string {
  return "VisualShaderNodeTexture3D"
}



// Enums

func (me *VisualShaderNodeTexture3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTexture3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeTexture3D) SetTexture(value Texture3D, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTexture3D) GetTexture()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
