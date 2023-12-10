// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func (me *VisualShaderNodeTexture2DArray) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTexture2DArray) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTexture2DArray) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeTexture2DArray) SetTextureArray(value Texture2DArray, )  {
  classNameV := StringNameFromStr("VisualShaderNodeTexture2DArray")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_array")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2206200446) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeTexture2DArray) GetTextureArray() Texture2DArray {
  classNameV := StringNameFromStr("VisualShaderNodeTexture2DArray")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_array")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 146117123) // FIXME: should cache?
  var ret Texture2DArray
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeTexture2DArray) GetPropTextureArray() Texture2DArray {
  panic("TODO: implement")
}

func (me *VisualShaderNodeTexture2DArray) SetPropTextureArray(value Texture2DArray) {
  panic("TODO: implement")
}