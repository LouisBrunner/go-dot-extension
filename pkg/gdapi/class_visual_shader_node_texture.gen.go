// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeTexture struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeTexture) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeTexture) BaseClass() string {
  return "VisualShaderNodeTexture"
}



// Enums

type VisualShaderNodeTextureSource int
const (
  VisualShaderNodeTextureSourceSourceTexture VisualShaderNodeTextureSource = 0
  VisualShaderNodeTextureSourceSourceScreen VisualShaderNodeTextureSource = 1
  VisualShaderNodeTextureSourceSource2DTexture VisualShaderNodeTextureSource = 2
  VisualShaderNodeTextureSourceSource2DNormal VisualShaderNodeTextureSource = 3
  VisualShaderNodeTextureSourceSourceDepth VisualShaderNodeTextureSource = 4
  VisualShaderNodeTextureSourceSourcePort VisualShaderNodeTextureSource = 5
  VisualShaderNodeTextureSourceSource3DNormal VisualShaderNodeTextureSource = 6
  VisualShaderNodeTextureSourceSourceRoughness VisualShaderNodeTextureSource = 7
  VisualShaderNodeTextureSourceSourceMax VisualShaderNodeTextureSource = 8
)

type VisualShaderNodeTextureTextureType int
const (
  VisualShaderNodeTextureTextureTypeTypeData VisualShaderNodeTextureTextureType = 0
  VisualShaderNodeTextureTextureTypeTypeColor VisualShaderNodeTextureTextureType = 1
  VisualShaderNodeTextureTextureTypeTypeNormalMap VisualShaderNodeTextureTextureType = 2
  VisualShaderNodeTextureTextureTypeTypeMax VisualShaderNodeTextureTextureType = 3
)

func (me *VisualShaderNodeTexture) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeTexture) SetSource(value VisualShaderNodeTextureSource, )  {
  classNameV := StringNameFromStr("VisualShaderNodeTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 905262939) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeTexture) GetSource() VisualShaderNodeTextureSource {
  classNameV := StringNameFromStr("VisualShaderNodeTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2896297444) // FIXME: should cache?
  var ret VisualShaderNodeTextureSource
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeTexture) SetTexture(value Texture2D, )  {
  classNameV := StringNameFromStr("VisualShaderNodeTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4051416890) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeTexture) GetTexture() Texture2D {
  classNameV := StringNameFromStr("VisualShaderNodeTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3635182373) // FIXME: should cache?
  var ret Texture2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeTexture) SetTextureType(value VisualShaderNodeTextureTextureType, )  {
  classNameV := StringNameFromStr("VisualShaderNodeTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 986314081) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeTexture) GetTextureType() VisualShaderNodeTextureTextureType {
  classNameV := StringNameFromStr("VisualShaderNodeTexture")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3290430153) // FIXME: should cache?
  var ret VisualShaderNodeTextureTextureType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeTexture) GetPropSource() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeTexture) SetPropSource(value int) {
  panic("TODO: implement")
}

func (me *VisualShaderNodeTexture) GetPropTexture() Texture2D {
  panic("TODO: implement")
}

func (me *VisualShaderNodeTexture) SetPropTexture(value Texture2D) {
  panic("TODO: implement")
}

func (me *VisualShaderNodeTexture) GetPropTextureType() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeTexture) SetPropTextureType(value int) {
  panic("TODO: implement")
}