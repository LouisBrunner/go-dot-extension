// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RDShaderSPIRV struct {
  Resource
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
// FIXME: can't seem to be able to use those from this side of the API

// Signals
