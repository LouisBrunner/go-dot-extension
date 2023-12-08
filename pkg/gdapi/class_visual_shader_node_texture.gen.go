// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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

type VisualShaderNodeTextureSource int
const (
  VisualShaderNodeTextureSourceTexture VisualShaderNodeTextureSource = 0
  VisualShaderNodeTextureSourceScreen VisualShaderNodeTextureSource = 1
  VisualShaderNodeTextureSource2DTexture VisualShaderNodeTextureSource = 2
  VisualShaderNodeTextureSource2DNormal VisualShaderNodeTextureSource = 3
  VisualShaderNodeTextureSourceDepth VisualShaderNodeTextureSource = 4
  VisualShaderNodeTextureSourcePort VisualShaderNodeTextureSource = 5
  VisualShaderNodeTextureSource3DNormal VisualShaderNodeTextureSource = 6
  VisualShaderNodeTextureSourceRoughness VisualShaderNodeTextureSource = 7
  VisualShaderNodeTextureSourceMax VisualShaderNodeTextureSource = 8
)

type VisualShaderNodeTextureTextureType int
const (
  VisualShaderNodeTextureTypeData VisualShaderNodeTextureTextureType = 0
  VisualShaderNodeTextureTypeColor VisualShaderNodeTextureTextureType = 1
  VisualShaderNodeTextureTypeNormalMap VisualShaderNodeTextureTextureType = 2
  VisualShaderNodeTextureTypeMax VisualShaderNodeTextureTextureType = 3
)
