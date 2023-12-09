// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeBooleanConstant struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeBooleanConstant) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeBooleanConstant) BaseClass() string {
  return "VisualShaderNodeBooleanConstant"
}



// Enums

func (me *VisualShaderNodeBooleanConstant) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeBooleanConstant) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeBooleanConstant) SetConstant(constant bool, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeBooleanConstant) GetConstant()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
