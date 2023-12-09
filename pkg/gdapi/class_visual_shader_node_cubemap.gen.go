// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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
  panic("TODO: implement")
}

func  (me *VisualShaderNodeCubemap) GetSource()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeCubemap) SetCubeMap(value Cubemap, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeCubemap) GetCubeMap()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeCubemap) SetTextureType(value VisualShaderNodeCubemapTextureType, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeCubemap) GetTextureType()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
