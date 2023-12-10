// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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
  classNameV := StringNameFromStr("VisualShaderNodeBillboard")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_billboard_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1227463289) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&billboard_type), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeBillboard) GetBillboardType() VisualShaderNodeBillboardBillboardType {
  classNameV := StringNameFromStr("VisualShaderNodeBillboard")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_billboard_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3724188517) // FIXME: should cache?
  var ret VisualShaderNodeBillboardBillboardType
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *VisualShaderNodeBillboard) SetKeepScaleEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("VisualShaderNodeBillboard")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_keep_scale_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *VisualShaderNodeBillboard) IsKeepScaleEnabled() bool {
  classNameV := StringNameFromStr("VisualShaderNodeBillboard")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_keep_scale_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *VisualShaderNodeBillboard) GetPropBillboardType() int {
  panic("TODO: implement")
}

func (me *VisualShaderNodeBillboard) SetPropBillboardType(value int) {
  panic("TODO: implement")
}

func (me *VisualShaderNodeBillboard) GetPropKeepScale() bool {
  panic("TODO: implement")
}

func (me *VisualShaderNodeBillboard) SetPropKeepScale(value bool) {
  panic("TODO: implement")
}