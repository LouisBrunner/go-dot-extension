// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeUIntConstant struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeUIntConstant) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeUIntConstant) BaseClass() string {
  return "VisualShaderNodeUIntConstant"
}



// Enums

func (me *VisualShaderNodeUIntConstant) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeUIntConstant) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeUIntConstant) SetConstant(constant int, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeUIntConstant) GetConstant()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
