// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeTexture2DArray struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeTexture2DArray) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeTexture2DArray) BaseClass() string {
  return "VisualShaderNodeTexture2DArray"
}



// Enums

func (me *VisualShaderNodeTexture2DArray) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTexture2DArray) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeTexture2DArray) SetTextureArray(value Texture2DArray, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTexture2DArray) GetTextureArray()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
