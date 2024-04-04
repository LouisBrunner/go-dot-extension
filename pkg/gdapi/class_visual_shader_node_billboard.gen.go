// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type VisualShaderNodeBillboard struct {
  VisualShaderNode
}

func (me *VisualShaderNodeBillboard) BaseClass() string {
  return "VisualShaderNodeBillboard"
}

func NewVisualShaderNodeBillboard() *VisualShaderNodeBillboard {
  str := StringNameFromStr("VisualShaderNodeBillboard") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &VisualShaderNodeBillboard{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&billboard_type) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeBillboard) GetBillboardType() VisualShaderNodeBillboardBillboardType {
  classNameV := StringNameFromStr("VisualShaderNodeBillboard")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_billboard_type")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3724188517) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  var ret VisualShaderNodeBillboardBillboardType

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(unsafe.Pointer(&ret)))
  return ret
}

func  (me *VisualShaderNodeBillboard) SetKeepScaleEnabled(enabled bool, )  {
  classNameV := StringNameFromStr("VisualShaderNodeBillboard")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_keep_scale_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *VisualShaderNodeBillboard) IsKeepScaleEnabled() bool {
  classNameV := StringNameFromStr("VisualShaderNodeBillboard")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_keep_scale_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
