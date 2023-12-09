// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeColorConstant struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeColorConstant) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeColorConstant) BaseClass() string {
  return "VisualShaderNodeColorConstant"
}



// Enums

func (me *VisualShaderNodeColorConstant) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeColorConstant) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeColorConstant) SetConstant(constant Color, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeColorConstant) GetConstant()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
