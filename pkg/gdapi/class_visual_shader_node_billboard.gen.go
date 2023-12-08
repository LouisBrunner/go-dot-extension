// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







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
  VisualShaderNodeBillboardBillboardTypeBillboardTypeDisabled VisualShaderNodeBillboardBillboardType = 0
  VisualShaderNodeBillboardBillboardTypeBillboardTypeEnabled VisualShaderNodeBillboardBillboardType = 1
  VisualShaderNodeBillboardBillboardTypeBillboardTypeFixedY VisualShaderNodeBillboardBillboardType = 2
  VisualShaderNodeBillboardBillboardTypeBillboardTypeParticles VisualShaderNodeBillboardBillboardType = 3
  VisualShaderNodeBillboardBillboardTypeBillboardTypeMax VisualShaderNodeBillboardBillboardType = 4
)

func  (me *VisualShaderNodeBillboard) SetBillboardType(billboard_type VisualShaderNodeBillboardBillboardType, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeBillboard) GetBillboardType() { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeBillboard) SetKeepScaleEnabled(enabled bool, ) { // TODO: return value
  // TODO: implement
}

func  (me *VisualShaderNodeBillboard) IsKeepScaleEnabled() { // TODO: return value
  // TODO: implement
}

// TODO: properties

// TODO: signals
