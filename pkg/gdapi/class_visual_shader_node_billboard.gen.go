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



// Enums

type VisualShaderNodeBillboardBillboardType int
const (
  VisualShaderNodeBillboardBillboardTypeBillboardTypeDisabled VisualShaderNodeBillboardBillboardType = 0
  VisualShaderNodeBillboardBillboardTypeBillboardTypeEnabled VisualShaderNodeBillboardBillboardType = 1
  VisualShaderNodeBillboardBillboardTypeBillboardTypeFixedY VisualShaderNodeBillboardBillboardType = 2
  VisualShaderNodeBillboardBillboardTypeBillboardTypeParticles VisualShaderNodeBillboardBillboardType = 3
  VisualShaderNodeBillboardBillboardTypeBillboardTypeMax VisualShaderNodeBillboardBillboardType = 4
)

func (me *VisualShaderNodeBillboard) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *VisualShaderNodeBillboard) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeBillboard) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeBillboard) SetBillboardType(billboard_type VisualShaderNodeBillboardBillboardType, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeBillboard) GetBillboardType()  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeBillboard) SetKeepScaleEnabled(enabled bool, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeBillboard) IsKeepScaleEnabled()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
