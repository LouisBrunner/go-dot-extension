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

func  (me *VisualShaderNodeCubemap) SetSource(value VisualShaderNodeCubemapSource, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeCubemap) GetSource() { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeCubemap) SetCubeMap(value Cubemap, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeCubemap) GetCubeMap() { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeCubemap) SetTextureType(value VisualShaderNodeCubemapTextureType, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeCubemap) GetTextureType() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
