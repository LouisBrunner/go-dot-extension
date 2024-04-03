// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RDShaderFile struct {
  Resource
}

func (me *RDShaderFile) BaseClass() string {
  return "RDShaderFile"
}

func NewRDShaderFile() *RDShaderFile {
  str := StringNameFromStr("RDShaderFile") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &RDShaderFile{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *RDShaderFile) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RDShaderFile) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RDShaderFile) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RDShaderFile) SetBytecode(bytecode RDShaderSPIRV, version StringName, )  {
  classNameV := StringNameFromStr("RDShaderFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_bytecode")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1558064255) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(bytecode.AsCTypePtr()), gdc.ConstTypePtr(version.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDShaderFile) GetSpirv(version StringName, ) RDShaderSPIRV {
  classNameV := StringNameFromStr("RDShaderFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_spirv")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3340165340) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(version.AsCTypePtr()), }
  ret := NewRDShaderSPIRV()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *RDShaderFile) GetVersionList() []StringName {
  classNameV := StringNameFromStr("RDShaderFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_version_list")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3995934104) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewArray()
  defer ret.Destroy()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ConvertArrayToSlice[StringName](ret)
}

func  (me *RDShaderFile) SetBaseError(error String, )  {
  classNameV := StringNameFromStr("RDShaderFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_base_error")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 83702148) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(error.AsCTypePtr()), }

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *RDShaderFile) GetBaseError() String {
  classNameV := StringNameFromStr("RDShaderFile")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_base_error")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 201670096) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  ret := NewString()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
