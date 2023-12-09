// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *VisualShaderNodeTexture) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeTexture) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeTexture) SetSource(value VisualShaderNodeTextureSource, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTexture) GetSource()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTexture) SetTexture(value Texture2D, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTexture) GetTexture()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTexture) SetTextureType(value VisualShaderNodeTextureTextureType, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeTexture) GetTextureType()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
