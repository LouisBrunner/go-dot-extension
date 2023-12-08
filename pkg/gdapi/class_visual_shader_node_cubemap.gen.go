// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
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

type VisualShaderNodeCubemapSource int
const (
  VisualShaderNodeCubemapSourceTexture VisualShaderNodeCubemapSource = 0
  VisualShaderNodeCubemapSourcePort VisualShaderNodeCubemapSource = 1
  VisualShaderNodeCubemapSourceMax VisualShaderNodeCubemapSource = 2
)

type VisualShaderNodeCubemapTextureType int
const (
  VisualShaderNodeCubemapTypeData VisualShaderNodeCubemapTextureType = 0
  VisualShaderNodeCubemapTypeColor VisualShaderNodeCubemapTextureType = 1
  VisualShaderNodeCubemapTypeNormalMap VisualShaderNodeCubemapTextureType = 2
  VisualShaderNodeCubemapTypeMax VisualShaderNodeCubemapTextureType = 3
)
