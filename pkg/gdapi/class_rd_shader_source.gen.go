// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RDShaderSource struct {
  obj gdc.ObjectPtr
}

func (me *RDShaderSource) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RDShaderSource) BaseClass() string {
  return "RDShaderSource"
}



// Enums

func (me *RDShaderSource) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RDShaderSource) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RDShaderSource) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RDShaderSource) SetStageSource(stage RenderingDeviceShaderStage, source String, )  {
  classNameV := StringNameFromStr("RDShaderSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stage_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 620821314) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stage), gdc.ConstTypePtr(source.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDShaderSource) GetStageSource(stage RenderingDeviceShaderStage, ) String {
  classNameV := StringNameFromStr("RDShaderSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stage_source")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3354920045) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stage), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDShaderSource) SetLanguage(language RenderingDeviceShaderLanguage, )  {
  classNameV := StringNameFromStr("RDShaderSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3422186742) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&language), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDShaderSource) GetLanguage() RenderingDeviceShaderLanguage {
  classNameV := StringNameFromStr("RDShaderSource")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_language")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1063538261) // FIXME: should cache?
  var ret RenderingDeviceShaderLanguage
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *RDShaderSource) GetPropSourceVertex() String {
  panic("TODO: implement")
}

func (me *RDShaderSource) SetPropSourceVertex(value String) {
  panic("TODO: implement")
}

func (me *RDShaderSource) GetPropSourceFragment() String {
  panic("TODO: implement")
}

func (me *RDShaderSource) SetPropSourceFragment(value String) {
  panic("TODO: implement")
}

func (me *RDShaderSource) GetPropSourceTesselationControl() String {
  panic("TODO: implement")
}

func (me *RDShaderSource) SetPropSourceTesselationControl(value String) {
  panic("TODO: implement")
}

func (me *RDShaderSource) GetPropSourceTesselationEvaluation() String {
  panic("TODO: implement")
}

func (me *RDShaderSource) SetPropSourceTesselationEvaluation(value String) {
  panic("TODO: implement")
}

func (me *RDShaderSource) GetPropSourceCompute() String {
  panic("TODO: implement")
}

func (me *RDShaderSource) SetPropSourceCompute(value String) {
  panic("TODO: implement")
}

func (me *RDShaderSource) GetPropLanguage() int {
  panic("TODO: implement")
}

func (me *RDShaderSource) SetPropLanguage(value int) {
  panic("TODO: implement")
}