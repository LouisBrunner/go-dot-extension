// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *RDShaderSPIRV) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RDShaderSPIRV) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *RDShaderSPIRV) SetStageBytecode(stage RenderingDeviceShaderStage, bytecode PackedByteArray, )  {
  panic("TODO: implement")
}

func  (me *RDShaderSPIRV) GetStageBytecode(stage RenderingDeviceShaderStage, )  {
  panic("TODO: implement")
}

func  (me *RDShaderSPIRV) SetStageCompileError(stage RenderingDeviceShaderStage, compile_error String, )  {
  panic("TODO: implement")
}

func  (me *RDShaderSPIRV) GetStageCompileError(stage RenderingDeviceShaderStage, )  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
