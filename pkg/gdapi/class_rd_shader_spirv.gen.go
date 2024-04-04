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

type RDShaderSPIRV struct {
  Resource
}

func (me *RDShaderSPIRV) BaseClass() string {
  return "RDShaderSPIRV"
}

func NewRDShaderSPIRV() *RDShaderSPIRV {
  str := StringNameFromStr("RDShaderSPIRV") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RDShaderSPIRV{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stage) , bytecode.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDShaderSPIRV) GetStageBytecode(stage RenderingDeviceShaderStage, ) PackedByteArray {
  classNameV := StringNameFromStr("RDShaderSPIRV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stage_bytecode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3816765404) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stage) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedByteArray()
  pinner.Pin(&stage)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RDShaderSPIRV) SetStageCompileError(stage RenderingDeviceShaderStage, compile_error String, )  {
  classNameV := StringNameFromStr("RDShaderSPIRV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stage_compile_error")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 620821314) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stage) , compile_error.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDShaderSPIRV) GetStageCompileError(stage RenderingDeviceShaderStage, ) String {
  classNameV := StringNameFromStr("RDShaderSPIRV")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stage_compile_error")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3354920045) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&stage) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewString()
  pinner.Pin(&stage)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
