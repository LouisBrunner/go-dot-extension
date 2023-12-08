// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeSample3D struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeSample3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeSample3D) BaseClass() string {
  return "VisualShaderNodeSample3D"
}

type VisualShaderNodeSample3DSource int
const (
  VisualShaderNodeSample3DSourceTexture VisualShaderNodeSample3DSource = 0
  VisualShaderNodeSample3DSourcePort VisualShaderNodeSample3DSource = 1
  VisualShaderNodeSample3DSourceMax VisualShaderNodeSample3DSource = 2
)
