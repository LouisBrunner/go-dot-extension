// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type VisualShaderNodeExpression struct {
  obj gdc.ObjectPtr
}

func (me *VisualShaderNodeExpression) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *VisualShaderNodeExpression) BaseClass() string {
  return "VisualShaderNodeExpression"
}



// Enums

func (me *VisualShaderNodeExpression) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *VisualShaderNodeExpression) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *VisualShaderNodeExpression) SetExpression(expression String, )  {
  panic("TODO: implement")
}

func  (me *VisualShaderNodeExpression) GetExpression()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
