// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type RDShaderFile struct {
  obj gdc.ObjectPtr
}

func (me *RDShaderFile) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RDShaderFile) BaseClass() string {
  return "RDShaderFile"
}



// Enums

func (me *RDShaderFile) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RDShaderFile) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *RDShaderFile) SetBytecode(bytecode RDShaderSPIRV, version StringName, )  {
  panic("TODO: implement")
}

func  (me *RDShaderFile) GetSpirv(version StringName, )  {
  panic("TODO: implement")
}

func  (me *RDShaderFile) GetVersionList()  {
  panic("TODO: implement")
}

func  (me *RDShaderFile) SetBaseError(error String, )  {
  panic("TODO: implement")
}

func  (me *RDShaderFile) GetBaseError()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
