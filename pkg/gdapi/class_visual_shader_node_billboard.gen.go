// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeBillboard struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeBillboard) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeBillboard) BaseClass() string {
  return "VisualShaderNodeBillboard"
}

type VisualShaderNodeBillboardBillboardType int
const (
  VisualShaderNodeBillboardBillboardTypeDisabled VisualShaderNodeBillboardBillboardType = 0
  VisualShaderNodeBillboardBillboardTypeEnabled VisualShaderNodeBillboardBillboardType = 1
  VisualShaderNodeBillboardBillboardTypeFixedY VisualShaderNodeBillboardBillboardType = 2
  VisualShaderNodeBillboardBillboardTypeParticles VisualShaderNodeBillboardBillboardType = 3
  VisualShaderNodeBillboardBillboardTypeMax VisualShaderNodeBillboardBillboardType = 4
)
