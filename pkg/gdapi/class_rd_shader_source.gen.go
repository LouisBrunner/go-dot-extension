// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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
  panic("TODO: implement")
}

func  (me *RDShaderSource) GetStageSource(stage RenderingDeviceShaderStage, )  {
  panic("TODO: implement")
}

func  (me *RDShaderSource) SetLanguage(language RenderingDeviceShaderLanguage, )  {
  panic("TODO: implement")
}

func  (me *RDShaderSource) GetLanguage()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
