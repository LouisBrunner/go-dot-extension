// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type VisualShaderNodeCubemap struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeCubemap) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeCubemap) BaseClass() string {
  return "VisualShaderNodeCubemap"
}



// Enums

type VisualShaderNodeCubemapSource int
const (
  VisualShaderNodeCubemapSourceSourceTexture VisualShaderNodeCubemapSource = 0
  VisualShaderNodeCubemapSourceSourcePort VisualShaderNodeCubemapSource = 1
  VisualShaderNodeCubemapSourceSourceMax VisualShaderNodeCubemapSource = 2
)

type VisualShaderNodeCubemapTextureType int
const (
  VisualShaderNodeCubemapTextureTypeTypeData VisualShaderNodeCubemapTextureType = 0
  VisualShaderNodeCubemapTextureTypeTypeColor VisualShaderNodeCubemapTextureType = 1
  VisualShaderNodeCubemapTextureTypeTypeNormalMap VisualShaderNodeCubemapTextureType = 2
  VisualShaderNodeCubemapTextureTypeTypeMax VisualShaderNodeCubemapTextureType = 3
)

func (me *VisualShaderNodeCubemap) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeCubemap) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeCubemap) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *VisualShaderNodeCubemap) SetSource(value VisualShaderNodeCubemapSource, )  {
  classNameV := StringNameFromStr("VisualShaderNodeCubemap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1625400621) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeCubemap) GetSource() VisualShaderNodeCubemapSource {
  classNameV := StringNameFromStr("VisualShaderNodeCubemap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2222048781) // FIXME: should cache?
  var ret VisualShaderNodeCubemapSource
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeCubemap) SetCubeMap(value Cubemap, )  {
  classNameV := StringNameFromStr("VisualShaderNodeCubemap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_cube_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2219800736) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(value.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeCubemap) GetCubeMap() Cubemap {
  classNameV := StringNameFromStr("VisualShaderNodeCubemap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_cube_map")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1772111058) // FIXME: should cache?
  var ret Cubemap
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeCubemap) SetTextureType(value VisualShaderNodeCubemapTextureType, )  {
  classNameV := StringNameFromStr("VisualShaderNodeCubemap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_texture_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1899718876) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&value), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeCubemap) GetTextureType() VisualShaderNodeCubemapTextureType {
  classNameV := StringNameFromStr("VisualShaderNodeCubemap")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_texture_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3356498888) // FIXME: should cache?
  var ret VisualShaderNodeCubemapTextureType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeCubemap) GetPropSource() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeCubemap) SetPropSource(value int) {
  panic("TODO: implement")
}

func (me *VisualShaderNodeCubemap) GetPropCubeMap() Cubemap {
  panic("TODO: implement")
}

func (me *VisualShaderNodeCubemap) SetPropCubeMap(value Cubemap) {
  panic("TODO: implement")
}

func (me *VisualShaderNodeCubemap) GetPropTextureType() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeCubemap) SetPropTextureType(value int) {
  panic("TODO: implement")
}