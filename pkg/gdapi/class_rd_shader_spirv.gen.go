// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RDShaderSPIRV struct {
  obj gdc.ObjectPtr
}

func (me *RDShaderSPIRV) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RDShaderSPIRV) BaseClass() string {
  return "RDShaderSPIRV"
}



// Enums

func (me *RDShaderSPIRV) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RDShaderSPIRV) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RDShaderSPIRV) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RDShaderSPIRV) SetStageBytecode(stage RenderingDeviceShaderStage, bytecode PackedByteArray, )  {
  classNameV := StringNameFromStr("RDShaderSPIRV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stage_bytecode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3514097977) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stage), gdc.ConstTypePtr(bytecode.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDShaderSPIRV) GetStageBytecode(stage RenderingDeviceShaderStage, ) PackedByteArray {
  classNameV := StringNameFromStr("RDShaderSPIRV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stage_bytecode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3816765404) // FIXME: should cache?
  var ret PackedByteArray
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stage), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RDShaderSPIRV) SetStageCompileError(stage RenderingDeviceShaderStage, compile_error String, )  {
  classNameV := StringNameFromStr("RDShaderSPIRV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stage_compile_error")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 620821314) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stage), gdc.ConstTypePtr(compile_error.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *RDShaderSPIRV) GetStageCompileError(stage RenderingDeviceShaderStage, ) String {
  classNameV := StringNameFromStr("RDShaderSPIRV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stage_compile_error")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3354920045) // FIXME: should cache?
  var ret String
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stage), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *RDShaderSPIRV) GetPropBytecodeVertex() PackedByteArray {
  panic("TODO: implement")
}

func (me *RDShaderSPIRV) SetPropBytecodeVertex(value PackedByteArray) {
  panic("TODO: implement")
}

func (me *RDShaderSPIRV) GetPropBytecodeFragment() PackedByteArray {
  panic("TODO: implement")
}

func (me *RDShaderSPIRV) SetPropBytecodeFragment(value PackedByteArray) {
  panic("TODO: implement")
}

func (me *RDShaderSPIRV) GetPropBytecodeTesselationControl() PackedByteArray {
  panic("TODO: implement")
}

func (me *RDShaderSPIRV) SetPropBytecodeTesselationControl(value PackedByteArray) {
  panic("TODO: implement")
}

func (me *RDShaderSPIRV) GetPropBytecodeTesselationEvaluation() PackedByteArray {
  panic("TODO: implement")
}

func (me *RDShaderSPIRV) SetPropBytecodeTesselationEvaluation(value PackedByteArray) {
  panic("TODO: implement")
}

func (me *RDShaderSPIRV) GetPropBytecodeCompute() PackedByteArray {
  panic("TODO: implement")
}

func (me *RDShaderSPIRV) SetPropBytecodeCompute(value PackedByteArray) {
  panic("TODO: implement")
}

func (me *RDShaderSPIRV) GetPropCompileErrorVertex() String {
  panic("TODO: implement")
}

func (me *RDShaderSPIRV) SetPropCompileErrorVertex(value String) {
  panic("TODO: implement")
}

func (me *RDShaderSPIRV) GetPropCompileErrorFragment() String {
  panic("TODO: implement")
}

func (me *RDShaderSPIRV) SetPropCompileErrorFragment(value String) {
  panic("TODO: implement")
}

func (me *RDShaderSPIRV) GetPropCompileErrorTesselationControl() String {
  panic("TODO: implement")
}

func (me *RDShaderSPIRV) SetPropCompileErrorTesselationControl(value String) {
  panic("TODO: implement")
}

func (me *RDShaderSPIRV) GetPropCompileErrorTesselationEvaluation() String {
  panic("TODO: implement")
}

func (me *RDShaderSPIRV) SetPropCompileErrorTesselationEvaluation(value String) {
  panic("TODO: implement")
}

func (me *RDShaderSPIRV) GetPropCompileErrorCompute() String {
  panic("TODO: implement")
}

func (me *RDShaderSPIRV) SetPropCompileErrorCompute(value String) {
  panic("TODO: implement")
}